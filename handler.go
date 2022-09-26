package main

import (
	"fmt"
	"github.com/Ccc-me/for-golang-test/db/mongodb"
	"github.com/Ccc-me/for-golang-test/db/mysql"
	"github.com/Ccc-me/for-golang-test/db/redis"
	"io/ioutil"
	"net/http"
	"os"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	os.Stdout.WriteString("Msg to hello\n")
	fmt.Fprintf(w, "hello\n")
}

func Headers(w http.ResponseWriter, req *http.Request) {
	os.Stdout.WriteString("Msg to headers\n")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func Err1(w http.ResponseWriter, req *http.Request) {
	os.Stderr.WriteString("Msg to err1\n")
	http.Error(w, "this is an err interface", 500)
}

func Err2(w http.ResponseWriter, req *http.Request) {
	os.Stderr.WriteString("Msg to err2\n")
	http.Error(w, "this is an err interface", 404)
}

func Ping(w http.ResponseWriter, req *http.Request) {
	os.Stdout.WriteString("Msg to ping\n")
	fmt.Fprintf(w, "pong!\n")
}

func Body(w http.ResponseWriter, req *http.Request) {
	b, e := ioutil.ReadAll(req.Body)
	fmt.Fprintf(w, "body: %v, err: %v", string(b), e)
}

func TestPanic(w http.ResponseWriter, req *http.Request) {
	os.Stderr.WriteString("Msg to testPanic\n")
	panic(req)
}

func Log(w http.ResponseWriter, req *http.Request) {
	os.Stdout.WriteString("FATAL 1658217911838250000 example.go:66 10.79.163.90 fatal level test!\n")
	os.Stdout.WriteString("WARN 1658217911838250001 example.go:66 10.79.163.90 warn level test!\n")
	os.Stdout.WriteString("ERROR 1658217911838250002 example.go:66 10.79.163.90 error level test!\n")
	os.Stdout.WriteString("NOTICE 1658217911838250002 example.go:66 10.79.163.90 notice level test!\n")
	os.Stdout.WriteString("INFO 1658217911838250002 example.go:66 10.79.163.90 info level test!\n")
	os.Stdout.WriteString("DEBUG 1658217911838250002 example.go:66 10.79.163.90 debug level test!\n")
}

func MysqlSelect(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	model, err := mysql.Select(id)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", model, err)
}

func MysqlCreate(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	model, err := mysql.Create(name)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", model, err)
}

func MysqlCreateLockTable(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	model, err := mysql.CreateLockTable(name)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", model, err)
}

func MysqlUpdate(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	model, err := mysql.Update(id)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", model, err)
}

func MysqlUpdateCounts(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	err := mysql.UpdateCounts(name)
	fmt.Fprintf(w, "err: %v\n", err)
}


func MysqlDelete(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	err := mysql.Delete(id)
	fmt.Fprintf(w, "err: %v\n", err)
}

func RedisSet(w http.ResponseWriter, req *http.Request) {
	key := req.FormValue("key")
	value := req.FormValue("value")
	expireTime := req.FormValue("expireTime")
	res, err := redis.Set(key, value, expireTime)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", res, err)
}

func RedisGet(w http.ResponseWriter, req *http.Request) {
	key := req.FormValue("key")
	res, err := redis.Get(key)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", res, err)
}

func RedisDel(w http.ResponseWriter, req *http.Request) {
	key := req.FormValue("key")
	res, err := redis.Del(key)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", res, err)
}

func MongoInsert(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	count := req.FormValue("count")
	res, err := mongodb.InsertOne(name, count)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", res, err)
}

func MongoFind(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	count := req.FormValue("count")
	res, err := mongodb.FindOne(name, count)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", res, err)
}

func MongoDelete(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	count := req.FormValue("count")
	res, err := mongodb.DeleteOne(name, count)
	fmt.Fprintf(w, "Response: %+v, err: %v\n", res, err)
}