package main

import (
	"fmt"
	"net/http"
)

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	http.ListenAndServe(":8080", nil)
}
func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `welcome to my first golang web page`)

}

func about(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `welcome to my first golang web page- about`)

}

func contact(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `welcome to my first golang web page- contact`)

}
