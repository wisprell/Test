package main

import (
	"github.com/Ccc-me/for-golang-test/db/mysql"
	"net/http"
)

func Init() {
	mysql.InitMysql()
}

func main() {
	Init()

	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/headers", Headers)
	http.HandleFunc("/v1/ping", Ping)
	http.HandleFunc("/err/500", Err1)
	http.HandleFunc("/err/404", Err2)
	http.HandleFunc("/vi/body", Body)
	http.HandleFunc("/panic", TestPanic)
	http.HandleFunc("/log", Log)

	http.HandleFunc("/mysql/select", MysqlSelect)
	http.HandleFunc("/mysql/create", MysqlCreate)
	http.HandleFunc("/mysql/update", MysqlUpdate)
	http.HandleFunc("/mysql/delete", MysqlDelete)

	http.ListenAndServe(":8000", nil)
}