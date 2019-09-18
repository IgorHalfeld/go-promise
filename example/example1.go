package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/igorhalfeld/go-promise"
)

func main() {
	p := promise.New(func(resolve chan interface{}, reject chan error) {
		resp, err := http.Get("http://gobyexample.com")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			resolve <- resp.Status
		} else {
			reject <- errors.New("ERRR")
		}
	})

	p.Then(func(value interface{}) {
		fmt.Println("Success", value)
	}).Catch(func(err error) {
		fmt.Println("Error", err)
	})

	p.Wait()
}
