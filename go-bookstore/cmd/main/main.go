package main

import (
	"github.com/gorilla/mux"
	"github.com/s-chance/go-bookstore/pkg/route"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	route.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
