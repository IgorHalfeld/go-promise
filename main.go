package main

import (
	"fmt"
	"net/http"

	"github.com/igorhalfeld/go-promise/promise"
)

func main() {
	request := promise.NewPromise(func(resolve chan interface{}, _ chan error) {
		resp, err := http.Get("http://gobyexample.com")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		resolve <- resp.Status
		// reject <- errors.New("Deu ruim")

	})

	request.Then(func(value interface{}, done func()) {
		fmt.Println("Success", value)
		done()
	})
	request.Catch(func(err error, done func()) {
		fmt.Println("Error", err)
		done()
	})
}
