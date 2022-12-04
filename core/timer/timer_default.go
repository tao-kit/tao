package timer

import (
	"github.com/sllt/tao/core/af"
	"sync/atomic"
	"time"
)

func New(options ...TimerOptions) *Timer {
	t := &Timer{
		queue:  newPriorityQueue(),
		status: StatusRunning,
		ticks:  0,
	}
	if len(options) > 0 {
		t.options = options[0]
	} else {
		t.options = DefaultOptions()
	}
	_ = af.Submit(func() {
		t.loop()
	})
	return t
}

// Add adds a timing job to the timer, which runs in interval of <interval>.
func (t *Timer) Add(interval time.Duration, job JobFunc) *Job {
	return t.createJob(interval, job, false, defaultTimes, StatusReady)
}

// AddJob adds a timing job to the timer with detailed parameters.
//
// The parameter <interval> specifies the running interval of the job.
//
// The parameter <singleton> specifies whether the job running in singleton mode.
// There's only one of the same job is allowed running when its a singleton mode job.
//
// The parameter <times> specifies limit for the job running times, which means the job
// exits if its run times exceeds the <times>.
//
// The parameter <status> specifies the job status when it's firstly added to the timer.
func (t *Timer) AddJob(interval time.Duration, job JobFunc, singleton bool, times int, status int) *Job {
	return t.createJob(interval, job, singleton, times, status)
}

// AddSingleton is a convenience function for add singleton mode job.
func (t *Timer) AddSingleton(interval time.Duration, job JobFunc) *Job {
	return t.createJob(interval, job, true, defaultTimes, StatusReady)
}

// AddOnce is a convenience function for adding a job which only runs once and then exits.
func (t *Timer) AddOnce(interval time.Duration, job JobFunc) *Job {
	return t.createJob(interval, job, true, 1, StatusReady)
}

// AddTimes is a convenience function for adding a job which is limited running times.
func (t *Timer) AddTimes(interval time.Duration, times int, job JobFunc) *Job {
	return t.createJob(interval, job, true, times, StatusReady)
}

// DelayAdd adds a timing job after delay of <interval> duration.
// Also see Add.
func (t *Timer) DelayAdd(delay time.Duration, interval time.Duration, job JobFunc) {
	t.AddOnce(delay, func() {
		t.Add(interval, job)
	})
}

// DelayAddJob adds a timing job after delay of <interval> duration.
// Also see AddJob.
func (t *Timer) DelayAddJob(delay time.Duration, interval time.Duration, job JobFunc, singleton bool, times int, status int) {
	t.AddOnce(delay, func() {
		t.AddJob(interval, job, singleton, times, status)
	})
}

// DelayAddSingleton adds a timing job after delay of <interval> duration.
// Also see AddSingleton.
func (t *Timer) DelayAddSingleton(delay time.Duration, interval time.Duration, job JobFunc) {
	t.AddOnce(delay, func() {
		t.AddSingleton(interval, job)
	})
}

// DelayAddOnce adds a timing job after delay of <interval> duration.
// Also see AddOnce.
func (t *Timer) DelayAddOnce(delay time.Duration, interval time.Duration, job JobFunc) {
	t.AddOnce(delay, func() {
		t.AddOnce(interval, job)
	})
}

// DelayAddTimes adds a timing job after delay of <interval> duration.
// Also see AddTimes.
func (t *Timer) DelayAddTimes(delay time.Duration, interval time.Duration, times int, job JobFunc) {
	t.AddOnce(delay, func() {
		t.AddTimes(interval, times, job)
	})
}

// Start starts the timer.
func (t *Timer) Start() {
	atomic.SwapInt64(&t.status, StatusRunning)
}

// Stop stops the timer.
func (t *Timer) Stop() {
	atomic.SwapInt64(&t.status, StatusStopped)
}

// Close closes the timer.
func (t *Timer) Close() {
	atomic.SwapInt64(&t.status, StatusClosed)
}

// createJob creates and adds a timing job to the timer.
func (t *Timer) createJob(interval time.Duration, job JobFunc, singleton bool, times int, status int) *Job {
	if times <= 0 {
		times = defaultTimes
	}
	var (
		intervalTicksOfJob = int64(interval / t.options.Interval)
	)
	if intervalTicksOfJob == 0 {
		// If the given interval is lesser than the one of the wheel,
		// then sets it to one tick, which means it will be run in one interval.
		intervalTicksOfJob = 1
	}
	nextTicks := t.ticks + intervalTicksOfJob
	j := &Job{
		job:       job,
		timer:     t,
		ticks:     intervalTicksOfJob,
		times:     int64(times),
		status:    int64(status),
		singleton: singleton,
		nextTicks: nextTicks,
	}
	t.queue.Push(j, nextTicks)
	return j
}

// loop starts the ticker using a standalone goroutine.
func (t *Timer) loop() {
	_ = af.Submit(func() {
		var (
			currentTimerTicks   int64
			timerIntervalTicker = time.NewTicker(t.options.Interval)
		)
		defer timerIntervalTicker.Stop()
		for {
			select {
			case <-timerIntervalTicker.C:
				// Check the timer status.
				switch t.status {
				case StatusRunning:
					// Timer proceeding.
					currentTimerTicks = atomic.AddInt64(&t.ticks, 1)
					if currentTimerTicks >= t.queue.LatestPriority() {
						t.proceed(currentTimerTicks)
					}

				case StatusStopped:
					// Do nothing.

				case StatusClosed:
					// Timer exits.
					return
				}
			}
		}
	})
}

// proceed proceeds the timer job checking and running logic.
func (t *Timer) proceed(currentTimerTicks int64) {
	var (
		value interface{}
	)
	for {
		value = t.queue.Pop()
		if value == nil {
			break
		}
		job := value.(*Job)
		// It checks if it meets the ticks requirement.
		if jobNextTicks := job.nextTicks; currentTimerTicks < jobNextTicks {
			// It push the job back if current ticks does not meet its running ticks requirement.
			t.queue.Push(job, job.nextTicks)
			break
		}
		// It checks the job running requirements and then does asynchronous running.
		job.doCheckAndRunByTicks(currentTimerTicks)
		// Status check: push back or ignore it.
		if job.Status() != StatusClosed {
			// It pushes the job back to queue for next running.
			t.queue.Push(job, job.nextTicks)
		}
	}
}
