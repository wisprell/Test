package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func err(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "this is an err interface", 500)
}

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong!\n")
}

func main() {
	initTimer()
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/v1/ping", ping)
	http.HandleFunc("/err", err)

	http.ListenAndServe(":8000", nil)
}

func initTimer() {
	timer1 := time.NewTimer(5 * time.Second)
	fmt.Println("开始时间: ", time.Now().Format("2006-01-02 15:04:05"))

	go func() {
		for {
			<-timer1.C
			fmt.Println("timer", time.Now().Format("2006-01-02 15:04:05"))
			os.Stdout.WriteString("Msg to STDOUT\n")
			os.Stderr.WriteString("Msg to STDERR\n")
			timer1.Reset(10*time.Second)
		}
	}()
	time.Sleep(15 * time.Second)
	fmt.Println("结束时间：", time.Now().Format("2006-01-02 15:04:05"))
}