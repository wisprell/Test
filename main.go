package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	os.Stdout.WriteString("Msg to hello\n")
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	os.Stdout.WriteString("Msg to headers\n")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func err1(w http.ResponseWriter, req *http.Request) {
	os.Stderr.WriteString("Msg to err1\n")
	http.Error(w, "this is an err interface", 500)
}

func err2(w http.ResponseWriter, req *http.Request) {
	os.Stderr.WriteString("Msg to err2\n")
	http.Error(w, "this is an err interface", 404)
}

func ping(w http.ResponseWriter, req *http.Request) {
	os.Stdout.WriteString("Msg to ping\n")
	fmt.Fprintf(w, "pong!\n")
}

func body(w http.ResponseWriter, req *http.Request) {
	b, e := ioutil.ReadAll(req.Body)
	fmt.Fprintf(w, "body: %v, err: %v", string(b), e)
}

func testPanic(w http.ResponseWriter, req *http.Request) {
	os.Stderr.WriteString("Msg to testPanic\n")
	panic(req)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/v1/ping", ping)
	http.HandleFunc("/err/500", err1)
	http.HandleFunc("/err/404", err2)
	http.HandleFunc("/vi/body", body)
	http.HandleFunc("/panic", testPanic)

	http.ListenAndServe(":8000", nil)
}