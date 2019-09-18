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
	})

	request.Then(func(success chan interface{}) {
		fmt.Println("Success", <-success)
	})
	// .Catch(func(err chan error) {
	// 	fmt.Println("Error", <-err)
	// })

}
