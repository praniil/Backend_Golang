package packages

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func PrintHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()
	fmt.Println("server: hello handler started")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "hello sir\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}

	// for name, headers := range r.Header{
	// 	for _, h := range headers{
	// 		fmt.Fprintf(w, "%v: %v \n", name, h)
	// 	}
	// }

	// fmt.Println(r.Method)
	// fmt.Println("Hello ")

}
func PrintDescription(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	ctx := r.Context()
	fmt.Println("server: description handler started")
	defer fmt.Println("server: description handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "my description")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	} 
	// 	fmt.Println(r.Method)
	// 	fmt.Println("My name is Pranil Parajuli and I am awesome")
}
