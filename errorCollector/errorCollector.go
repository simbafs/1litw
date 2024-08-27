package errorCollector

import (
	"errors"
	"sync"
)

// ErrorCollector collects errors thread-safely.
type ErrorCollector struct {
	errs []error
	lock sync.Mutex
}

func New() *ErrorCollector {
	return &ErrorCollector{
		errs: []error{},
		lock: sync.Mutex{},
	}
}

func (ec *ErrorCollector) Add(err error) {
	ec.lock.Lock()
	defer ec.lock.Unlock()

	if err == nil {
		return
	}

	ec.errs = append(ec.errs, err)
}

func (ec *ErrorCollector) Error() string {
	return errors.Join(ec.errs...).Error()
}

func (ec *ErrorCollector) IsEmpty() bool {
	return len(ec.errs) == 0
}
