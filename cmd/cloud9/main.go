package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

const PORT = ":8080"

func main(){
	router := mux.NewRouter();

	router.HandleFunc("/health-check", healthCheck).Methods("GET")
	router.HandleFunc("/members", members).Methods("GET")

	membersRouter := router.PathPrefix("/members").Subrouter()
	
	membersRouter.HandleFunc("/{nickname}", memberByNickname).Methods("GET")
	membersRouter.HandleFunc("/name/{firstName}-{secondName}", memberByName).Methods("GET")
	membersRouter.HandleFunc("/type/{type}", membersByType).Methods("GET")

	log.Printf("Starting server on %s\n", PORT)
	err := http.ListenAndServe(PORT, router)
	log.Fatal(err)
}