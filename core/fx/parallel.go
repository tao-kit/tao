package fx

import (
	"github.com/tao-kit/tao/core/errorx"
	"github.com/tao-kit/tao/core/threading"
)

// Parallel runs fns parallelly and waits for done.
func Parallel(fns ...func()) {
	group := threading.NewRoutineGroup()
	for _, fn := range fns {
		group.RunSafe(fn)
	}
	group.Wait()
}

func ParallelErr(fns ...func() error) error {
	var be errorx.BatchError

	group := threading.NewRoutineGroup()
	for _, fn := range fns {
		f := fn
		group.RunSafe(func() {
			if err := f(); err != nil {
				be.Add(err)
			}
		})
	}
	group.Wait()

	return be.Err()
}
