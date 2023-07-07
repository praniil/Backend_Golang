package packages

import (
	"backend/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
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

}

type User struct {
	gorm.Model
	Name    string
	Surname string
}

func PrintDescription(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	db := database.Database_connection()

	ctx := r.Context()
	fmt.Println("server: description handler started")
	defer fmt.Println("server: description handler ended")

	select {
	case <-time.After(1 * time.Second):
		db.AutoMigrate(&User{})
		db.Create(&User{
			Name:    "pranil",
			Surname: "parajuli",
		})
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
	// 	fmt.Println(r.Method)
	// 	fmt.Println("My name is Pranil Parajuli and I am awesome")
}
