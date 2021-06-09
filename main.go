package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	//Declare the static file directory and point it to the
	//directory we just made
	staticFileDirectory := http.Dir("./assets")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	//The "pathPerfix" method acts as a matcher, and matches all routes starting
	//with "/assets", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
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
