package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Main")
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/movies", GetAllMovies).Methods("GET")
	r.HandleFunc("/api/v1/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/api/v1/movies", UpdateMovie).Methods("PUT")
	r.HandleFunc("/api/v1/movies", RemoveMovie).Methods("DELETE")
	r.HandleFunc("/api/v1/movies/{id}", GetMovie).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet")
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet")
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented yet")
}

func RemoveMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented yet")
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented yet")
}
