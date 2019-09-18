package promise

// Promise represents a promise struct
type Promise struct {
	success chan interface{}
	failure chan error
}

var resolve = make(chan interface{}, 1)
var reject = make(chan error, 1)

// NewPromise represents a new instance of Promise struct
func NewPromise(fn func(chan interface{}, chan error)) *Promise {
	promise := new(Promise)

	go fn(resolve, reject)

	go func() {
		select {
		case success := <-resolve:
			promise.success <- success
		case err := <-reject:
			promise.failure <- err
		}
	}()

	return promise
}

// Then represents a next chain of success flow
func (p Promise) Then(success func(chan interface{})) Promise {
	result := make(chan interface{}, 1)
	for {
		select {
		case result <- resolve:
			success(result)
		}
	}
	return p
}

// Catch represents a next chain of failure flow
// func (p *Promise) Catch(failure func(chan error)) *Promise {
// 	result := make(chan error, 1)
// 	select {
// 	case result <- p.failure:
// 		failure(result)
// 	}
// 	return p
// }
