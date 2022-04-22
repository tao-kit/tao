package timer

import (
	"manlu.org/tao/core/af"
	"math"
	"sync/atomic"
)

// Job is the timing job.
type Job struct {
	job       JobFunc // The job function.
	timer     *Timer  // Belonged timer.
	ticks     int64   // The job runs every ticks.
	times     int64   // Limit running times.
	status    int64   // Job status.
	singleton bool    // Singleton mode.
	nextTicks int64   // Next run ticks of the job.
}

// JobFunc is the job function.
type JobFunc = func()

// Status returns the status of the job.
func (j *Job) Status() int {
	return int(j.status)
}

// Run runs the timer job asynchronously.
func (j *Job) Run() {
	leftRunningTimes := atomic.AddInt64(&j.times, int64(-1))
	if leftRunningTimes < 0 {
		atomic.SwapInt64(&j.status, StatusClosed)
		return
	}
	// This means it does not limit the running times.
	// I know it's ugly, but it is surely high performance for running times limit.
	if leftRunningTimes < 2000000000 && leftRunningTimes > 1000000000 {
		atomic.SwapInt64(&j.times, math.MaxInt32)
	}

	_ = af.Submit(func() {
		defer func() {
			if err := recover(); err != nil {
				if err != panicExit {
					panic(err)
				} else {
					j.Close()
					return
				}
			}
			if j.Status() == StatusRunning {
				j.SetStatus(StatusReady)
			}
		}()
		j.job()
	})
}

// doCheckAndRunByTicks checks the if job can run in given timer ticks,
// it runs asynchronously if the given `currentTimerTicks` meets or else
// it increments its ticks and waits for next running check.
func (j *Job) doCheckAndRunByTicks(currentTimerTicks int64) {
	// Ticks check.
	if currentTimerTicks < j.nextTicks {
		return
	}
	atomic.SwapInt64(&j.nextTicks, currentTimerTicks+j.ticks)
	// Perform job checking.
	switch j.status {
	case StatusRunning:
		if j.IsSingleton() {
			return
		}
	case StatusReady:
		if !atomic.CompareAndSwapInt64(&j.status, StatusReady, StatusRunning) {
			return
		}
	case StatusStopped:
		return
	case StatusClosed:
		return
	}
	// Perform job running.
	j.Run()
}

// SetStatus custom sets the status for the job.
func (j *Job) SetStatus(status int) int {
	return int(atomic.SwapInt64(&j.status, int64(status)))
}

// Start starts the job.
func (j *Job) Start() {
	atomic.SwapInt64(&j.status, int64(StatusReady))
}

// Stop stops the job.
func (j *Job) Stop() {
	atomic.SwapInt64(&j.status, int64(StatusStopped))
}

// Close closes the job, and then it will be removed from the timer.
func (j *Job) Close() {
	atomic.SwapInt64(&j.status, int64(StatusClosed))
}

// Reset reset the job, which resets its ticks for next running.
func (j *Job) Reset() {
	atomic.SwapInt64(&j.nextTicks, j.timer.ticks+j.ticks)
}

// IsSingleton checks and returns whether the job in singleton mode.
func (j *Job) IsSingleton() bool {
	return j.singleton
}

// SetSingleton sets the job singleton mode.
func (j *Job) SetSingleton(enabled bool) {
	j.singleton = enabled
}

// Job returns the job function of this job.
func (j *Job) Job() JobFunc {
	return j.job
}

// SetTimes sets the limit running times for the job.
func (j *Job) SetTimes(times int) {
	atomic.SwapInt64(&j.times, int64(times))
}
