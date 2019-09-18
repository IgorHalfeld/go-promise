package main

import (
	"fmt"
	"net/http"

	"github.com/igorhalfeld/go-promise/promise"
)

func main() {
	p := promise.NewPromise(func(resolve chan interface{}, reject chan error) {
		resp, err := http.Get("http://gobyexample.com")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		resolve <- resp.Status
		// reject <- errors.New("Deu ruim")
	})

	p.Then(func(value interface{}) {
		fmt.Println("Success", value)
	}).Catch(func(err error) {
		fmt.Println("Error", err)
	})

	p.Wait()
}
