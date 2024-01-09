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

// structs - movies and director
type Anime struct{
	ID string `json: "id"`
	Title string `json: "title"`
	Desc string `json: "desc"`
	Studio *Studio `json: "director"`
}

type Studio struct{
	Firstname string `json: "firstname"`
	Lastname string `json: "lastname"`
}


func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello!!!")
	return
}

func getAnime(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	for ind, item := range animes{
		if item.ID == vars["id"]{
			json.NewEncoder(w).Encode(animes[ind])
		}
	}
	return
}

func getAnimes(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animes)
	return
}

func addAnime(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var newAnime Anime
	json.NewDecoder(r.Body).Decode(&newAnime)
	newAnime.ID = strconv.Itoa(rand.Intn(10000))
	animes = append(animes, newAnime)
	json.NewEncoder(w).Encode(animes)
	return
}

func updateAnime(w http.ResponseWriter, r *http.Request){
	 // set content type, params, loop over and delete and then add and then update w
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	fmt.Println(vars)
	for index, item := range animes{
		if item.ID == vars["ID"]{
			// delete annd then add
			animes = append(animes[:index], animes[index:]...)
			break
		}
	}
	var newAnime Anime
	json.NewDecoder(r.Body).Decode(&newAnime)
	newAnime.ID = vars["id"]
	json.NewEncoder(w).Encode(animes)
	return
}

func deleteAnime(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	for index, item := range animes{
		if item.ID == vars["id"]{
			// delete the element wiht ID == id
			animes = append(animes[:index], animes[index+1:]...)
			break
		}
	}
	return
}



var animes []Anime
func main(){
	r := mux.NewRouter()
	fmt.Println("Hello")
	r.HandleFunc("/", helloHandler)
	animes = append(animes, Anime{
		ID:        "1",
		Title:     "Your Name",
		Desc:      "A story about two high schoolers who mysteriously swap bodies and embark on a journey to unravel the truth behind their connection.",
		Studio:  &Studio{Firstname: "Makoto", Lastname:  "Shinkai"},
	})
	

	r.HandleFunc("/animes", getAnimes).Methods("GET")
	r.HandleFunc("/anime/{id}", getAnime).Methods("GET")
	r.HandleFunc("/anime", addAnime).Methods("POST")
	r.HandleFunc("/anime/{id}", updateAnime).Methods("PUT")
	r.HandleFunc("/anime/{id}", deleteAnime).Methods("DELETE")

	fmt.Println("Server running ")
	if err := http.ListenAndServe(":8080", r); err != nil{
		log.Fatal(err)
	}
}