package main
//for installing dependencies you need to have go.mod

import (
	"fmt" //for basic printing
	"log" //for logging errors
	"encoding/json" //to encode data into JSON while sharing on Postman
	"math/rand" //to create new id for new movie
	"net/http" //for creating HTTP server and handling requests
	"strconv" //to convert id into string
	"github.com/gorilla/mux" //for routing
)

type Movie struct {
	ID string `json: "id"`
	// The json tag is used to specify how the field should be encoded in JSON
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json: "lastname"`
}

var movies []Movie //slice of movies

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	//this tells the client that the response will be in JSON format
	json.NewEncoder(w).Encode(movies)
	//Encode encodes the movies slice into JSON and writes it to the response writer
}

func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//mux.Vars(r) returns a map of the variables extracted from the URL

	for index, item := range movies {

		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			// Remove the movie from the slice by slicing it before and after the index
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter,r *http.Request){
    w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie 
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// Decode the JSON request body into the movie variable
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	// Append the new movie to the movies slice
	json.NewEncoder(w).Encode(movie)
	// Encode the movie variable into JSON and write it to the response writer
}

func updateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode((&movie))
			movie.ID = params["id"]
			movies = append(movies, movie)
			// Append the updated movie to the movies slice
			json.NewEncoder(w).Encode(movie)
			// Encode the updated movie into JSON and write it to the response writer
			return
		}
	}
}

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438743", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45443", Title: "Movie Two", Director: &Director{Firstname: "Shirley", Lastname: "Setia"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting the server at port 8000...")
	log.Fatal(http.ListenAndServe(":8000",r))
	// ListenAndServe starts an HTTP server with a given address and handler which handles incoming requests.
}