package main

import (
	"fmt"
	"net/http"
	"time"
)

type Employee struct {
	firstName string
	lastName  string
}

func main() {
	println("starting server")

	employee := Employee{"vilas", "athavale"}

	var employee2 Employee
	employee2.firstName = "Jeff"
	employee2.lastName = "Prestes"

	for i := 0; i < 10; i++ {
		go runInParallel(employee)
		go runInParallel(employee2)
	}

	time.Sleep(4 * time.Second)

	fmt.Println("Employee is: %v", employee)
	fmt.Println("Employee is: %v", employee2)

	s := &http.Server{
		Addr:           ":8081",
		Handler:        handler1(),
		ReadTimeout:    600 * time.Second,
		WriteTimeout:   600 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

func runInParallel(e Employee) {
	time.Sleep(3 * time.Second)
	e.firstName = "bad"
	println("sleep over %s %s %p", e.firstName, e.lastName, &e)
}

func handler1() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		f, _ := w.(http.Flusher)
		defer f.Flush()

		reqType := req.FormValue("reqType")
		key := req.FormValue("key")
		value := req.FormValue("value")
		var v []byte

		if reqType == "get" {
			val := fmt.Sprintf("%s %s", "get", key)
			v = []byte(val)

		} else if reqType == "put" {
			val := fmt.Sprintf("%s %s", "get", value)
			v = []byte(val)
		} else if reqType == "remove" {
			val := fmt.Sprintf("%s %s", "get", key)
			v = []byte(val)
		}
		value = string(v)
		w.Write([]byte(value))
	}
	return http.HandlerFunc(fn)
}
