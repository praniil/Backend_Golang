package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", printHello)
	http.HandleFunc("/description", printDescription)
	//starting server
	log.Println("starting server local host:8080")
	http.ListenAndServe(":8080", nil)
}

func printHello(w http.ResponseWriter, r *http.Request) {

	
	if r.Method != "GET" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// for name, headers := range r.Header{
	// 	for _, h := range headers{
	// 		fmt.Fprintf(w, "%v: %v \n", name, h)
	// 	}
	// }

	fmt.Println(r.Method)
	fmt.Println("Hello ")

}

func printDescription(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	fmt.Println(r.Method)
	fmt.Println("My name is Pranil Parajuli and I am awesome")
}
