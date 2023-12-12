package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	msg := "Hello from Go, Chris is doing test"
	fmt.Fprintf(w, msg)

}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	fmt.Println("Application started and listening on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}
}
