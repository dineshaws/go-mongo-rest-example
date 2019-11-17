package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	. "github.com/dineshaws/go-mongo-rest-example/config"
	. "github.com/dineshaws/go-mongo-rest-example/dao"
	"github.com/gorilla/mux"
)

var config = Config{}
var dao = MoviesDAO{}

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

func init() {
	fmt.Println("init called")
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "not implemented yet")
	movies, err := dao.FindAll()
	if err != nil {
		fmt.Fprintln(w, err)
	}
	respondWithJson(w, http.StatusOK, movies)
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

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
