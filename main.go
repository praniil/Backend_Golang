package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", printHello)
	http.HandleFunc("/description", printDescription)
	//starting server
	log.Println("starting server local host:8080")
	http.ListenAndServe(":8080", nil)
}

func printHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")		//defer is used to delay the execution of a function until surrounding function completes

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "hello sir\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}

	if r.Method != "GET" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// for name, headers := range r.Header{
	// 	for _, h := range headers{
	// 		fmt.Fprintf(w, "%v: %v \n", name, h)
	// 	}
	// }

	// fmt.Println(r.Method)
	// fmt.Println("Hello ")

}

func printDescription(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("server: description handler started")
	defer fmt.Println("server: description handler ended")

	select{
	case <- time.After(10* time.Second):
		fmt.Fprintf(w, "my description")
	case <- ctx.Done():
		err:= ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	fmt.Println(r.Method)
	fmt.Println("My name is Pranil Parajuli and I am awesome")
}
