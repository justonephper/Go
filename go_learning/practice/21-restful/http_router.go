package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "Welcome\n")
}

func Hello(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	fmt.Fprintf(writer, "hello,%s", ps.ByName("name"))
}

func main() {
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(":8080", router))
}
