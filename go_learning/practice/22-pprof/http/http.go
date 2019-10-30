package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func GetFibonacciSerie(n int) []int {
	ret := make([]int, 2, n)
	ret[0] = 1
	ret[1] = 1
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "welcome")
}

func Hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello hello"))
}

func createFBS(w http.ResponseWriter, r *http.Request) {
	var fbs []int
	for i := 0; i < 1000000; i++ {
		fbs = GetFibonacciSerie(50)
	}
	w.Write([]byte(fmt.Sprintf("%v", fbs)))

}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/fb", createFBS)

	http.ListenAndServe(":8080", nil)
}
