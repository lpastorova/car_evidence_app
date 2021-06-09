package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}

func main() {
	r := newRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Err ListenAndServe", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fprintf, err := fmt.Fprintf(w, "hello")
	if err != nil {
		fmt.Println("Err inside handler:", err)
	}
	fmt.Println("/hello hit, with n =", fprintf)
}
