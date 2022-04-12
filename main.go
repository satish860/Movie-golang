package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
    FirstName string `json:"first_name"`
	LastName string `json:"last_name"`	
}

var movies []Movie;

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	params :=mux.Vars(r)
	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]... )
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	params :=mux.Vars(r)

	for _,item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	params :=mux.Vars(r)

	for index, item := range movies{

		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]... )
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main(){
	r := mux.NewRouter();
	movies = append(movies, Movie{ID: "1",Isbn: "323422",Title: "Hello Books",Director: &Director{FirstName: "john",LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2",Isbn: "3432444",Title: "Programming c#",Director: &Director{FirstName: "satish",LastName: "Venkatakrishnan"}})
	r.HandleFunc("/movies",getMovies).Methods("GET");
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovies).Methods("DELETE")
	fmt.Printf("Starting the server at 3000\n")
	log.Fatal(http.ListenAndServe(":3000",r))
}