package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/igorhalfeld/go-promise/promise"
)

func main() {
	request := promise.NewPromise(func(_ chan interface{}, reject chan error) {
		resp, err := http.Get("http://gobyexample.com")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// resolve <- resp.Status
		reject <- errors.New("Deu ruim")
	})

	request.Then(func(value interface{}) {
		fmt.Println("Success", value)
	})
	request.Catch(func(err error) {
		fmt.Println("Error", err)
	})

}
