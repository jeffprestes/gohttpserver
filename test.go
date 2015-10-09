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

	s := &http.Server{
		Addr:           ":8081",
		Handler:        handlerBootstrap4(),
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

func handlerBootstrap4() http.Handler {
	fn := func(writer http.ResponseWriter, req *http.Request) {

		writeHtmlHeader(writer)
		fmt.Fprintln(writer, "    <div class=\"container\">\n")
		fmt.Fprintln(writer, "      <h1>Test page: Golang & Bootstrap 4</h1>\n")
		fmt.Fprintln(writer, "    </div>\n")
		writeHtmlFooter(writer)

	}

	return http.HandlerFunc(fn)
}

func writeHtmlHeader(writer http.ResponseWriter) {
	fmt.Fprintln(writer, "<!DOCTYPE html>\n")
	fmt.Fprintln(writer, "<html lang=\"en\">")
	fmt.Fprintln(writer, "  <head>")
	fmt.Fprintln(writer, "    <title>Bootstrap & Golang Example Template</title>")
	fmt.Fprintln(writer, "    <meta charset=\"utf-8\">")
	fmt.Fprintln(writer, "    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">")
	fmt.Fprintln(writer, "    <link rel=\"stylesheet\" href=\"https://cdn.rawgit.com/twbs/bootstrap/v4-dev/dist/css/bootstrap.css\">")
	fmt.Fprintln(writer, "  </head>")
	fmt.Fprintln(writer, "  <body>\n")
}

func writeHtmlFooter(writer http.ResponseWriter) {
	fmt.Fprintln(writer, "    <script src=\"https://ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js\"></script>\n")
	fmt.Fprintln(writer, "    <script src=\"https://cdn.rawgit.com/twbs/bootstrap/v4-dev/dist/js/bootstrap.js\"></script>\n")
	fmt.Fprintln(writer, "  </body>\n")
	fmt.Fprintln(writer, "</html>")
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

func callParallelProcess() {
	/*
		employee := Employee{"vilas", "athavale"}

		var employee2 Employee
		employee2.firstName = "Jeff"
		employee2.lastName = "Prestes"

		for i := 0; i < 10; i++ {
			go runInParallel(employee)
			go runInParallel(employee2)
		}

		time.Sleep(4 * time.Second)

		fmt.Fprintln("Employee is: %v", employee)
		fmt.Fprintln("Employee is: %v", employee2)
	*/
}
