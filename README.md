## Go Promise 🐰

Manage async flow as a javascript person

### How to use

```go
package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/igorhalfeld/go-promise"
)

func main() {
	p := promise.New(func(resolve func(interface{}), reject func(error)) {
		resp, err := http.Get("http://gobyexample.com")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			resolve(resp.Status)
		} else {
			reject(errors.New("ERRR"))
		}
	})

	p.Then(func(value interface{}) {
		fmt.Println("Success", value)
	}).Catch(func(err error) {
		fmt.Println("Error", err)
	})

	p.Wait()
}

```

### Contributors

![brenin](https://avatars3.githubusercontent.com/u/16777941?s=100&v=4)  | ![João Pedro](https://avatars0.githubusercontent.com/u/4886125?s=100&v=4)
------------------------------------------------------------------------|------------------------------------------------------------------------ 
[Breno Andrade](https://github.com/BrenoAndrade)                        | [João Pedro](https://github.com/joaopmgd)