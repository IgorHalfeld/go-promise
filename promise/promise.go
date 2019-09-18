package promise

import "sync"

// Promise represents a promise struct
type Promise struct{}

var wg sync.WaitGroup
var resolve = make(chan interface{}, 1)
var reject = make(chan error, 1)

// NewPromise represents a new instance of Promise struct
func NewPromise(fn func(chan interface{}, chan error)) *Promise {
	promise := new(Promise)
	go fn(resolve, reject)
	return promise
}

// Then represents a next chain of success flow
func (p *Promise) Then(success func(interface{}, func())) *Promise {
	wg.Add(1)
	go func() {
		select {
		case result := <-resolve:
			success(result, done)
			wg.Done()
		}
	}()
	wg.Wait()
	return p
}

// Catch represents a next chain of failure flow
func (p *Promise) Catch(failure func(error, func())) *Promise {
	wg.Add(1)
	go func() {
		select {
		case result := <-reject:
			failure(result, done)
			wg.Done()
		}
	}()
	wg.Wait()
	return p
}

func done() {
	if _, ok := <-resolve; ok {
		close(resolve)
		wg.Done()
	}

	if _, ok := <-reject; ok {
		close(reject)
		wg.Done()
	}
}
