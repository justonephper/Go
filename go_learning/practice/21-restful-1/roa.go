package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"math/rand"
	"net/http"
	"strconv"
)

type student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var studentDB map[string]*student

func init() {
	studentDB = map[string]*student{}
	studentDB["mike"] = &student{"ID-0001", "mike", 25}
	studentDB["liming"] = &student{"ID-0002", "liming", 28}
}

func Index(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Fprint(writer, "welcome:"+strconv.Itoa(rand.Intn(100)))
}

func Hello(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	fmt.Fprintf(writer, "hello %s", ps.ByName("name"))
}

func GetStudentByName(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	var (
		ok       bool
		info     *student
		infoJson []byte
		err      error
	)
	if info, ok = studentDB[name]; !ok {
		writer.Write([]byte("{\"error\":\"not found\"}"))
		return
	}
	if infoJson, err = json.Marshal(info); err != nil {
		writer.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", err)))
		return
	}
	if byte_num, err := writer.Write(infoJson); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("response str len:%d", byte_num)
	}
}

func main() {
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/student/:name", GetStudentByName)
	http.ListenAndServe(":8080", router)
}
