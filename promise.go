package promise

// Promise represents a promise struct
type Promise struct{}

var resolve = make(chan interface{}, 1)
var reject = make(chan error, 1)
var done = make(chan bool, 1)

// New represents a new instance of Promise struct
func New(fn func(chan interface{}, chan error)) *Promise {
	promise := new(Promise)
	go fn(resolve, reject)
	return promise
}

// Then represents a next chain of success flow
func (p *Promise) Then(success func(interface{})) *Promise {
	go func() {
		if result, ok := <-resolve; ok {
			done <- ok
			success(result)
			close(reject)
		}
	}()
	return p
}

// Catch represents a next chain of failure flow
func (p *Promise) Catch(failure func(error)) *Promise {
	go func() {
		if result, ok := <-reject; ok {
			done <- ok
			failure(result)
			close(resolve)
		}
	}()
	return p
}

func (p *Promise) Wait() {
	<-done
}
