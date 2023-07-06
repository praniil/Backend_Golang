package main

import (
	"backend/packages"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", packages.PrintHello)
	http.HandleFunc("/description", packages.PrintDescription)
	//starting server
	log.Println("starting server local host:8080")
	http.ListenAndServe(":8080", nil)
}
