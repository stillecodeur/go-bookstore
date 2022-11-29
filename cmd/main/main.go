package main

import (
	"fmt"
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Running Book Store at 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
