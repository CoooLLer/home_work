package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) (errLimit error) {
	var wg sync.WaitGroup
	var errCount int64

	exitChan := make(chan struct{})
	tasksChan := make(chan Task)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(tasksChan chan Task, exitChan chan struct{}) {
			defer func() {
				wg.Done()
			}()
			for {
				select {
				case <-exitChan:
					return

				default:
					task, ok := <-tasksChan
					if !ok {
						return
					}
					if err := task(); err != nil {
						atomic.AddInt64(&errCount, 1)
					}
				}
			}
		}(tasksChan, exitChan)
	}

	for _, task := range tasks {
		if atomic.LoadInt64(&errCount) >= int64(m) {
			exitChan <- struct{}{}
			errLimit = ErrErrorsLimitExceeded
			break
		}

		tasksChan <- task
	}

	close(tasksChan)
	wg.Wait()
	return
}
