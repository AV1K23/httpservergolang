package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var Counter int

func GetHendler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Counter равен", strconv.Itoa(Counter))
	} else {
		fmt.Fprintln(w, "поддерживаеться только метод GET")
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		Counter++
		fmt.Fprintln(w, "counter увеличен на 1")
	} else {
		fmt.Fprintln(w, "поддерживаеться только метод POST")
	}
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "hello world")
}

func main() {

	r := http.NewServeMux()

	r.HandleFunc("/hello", HelloHandler)
	r.HandleFunc("/post", PostHandler)
	r.HandleFunc("/get", GetHendler)
	http.ListenAndServe(":8080", r)

}
