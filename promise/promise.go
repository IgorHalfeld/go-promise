package promise

import "fmt"

// Promise represents a promise struct
type Promise struct{}

var resolve = make(chan interface{}, 1)
var reject = make(chan error, 1)

// NewPromise represents a new instance of Promise struct
func NewPromise(fn func(chan interface{}, chan error)) *Promise {
	promise := new(Promise)
	go fn(resolve, reject)

	return promise
}

// Then represents a next chain of success flow
func (p Promise) Then(success func(interface{})) {
	for {
		select {
		case result := <-resolve:
			success(result)
			break
		}
	}
}

// Catch represents a next chain of failure flow
func (p *Promise) Catch(failure func(error)) {
	for {
		select {
		case result := <-reject:
			failure(result)
			fmt.Println("Reject", result)
			break
		}
	}
}
