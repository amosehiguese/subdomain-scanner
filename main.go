package main

import (
	"log"
	"net/http"

	"github.com/amosehiguese/subdscanner/api"
	"github.com/amosehiguese/subdscanner/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/api/v1/scan/{domain}", api.GetDomain).Methods("GET")

	log.Println("server running...")
	http.ListenAndServe(":8080", r)
}