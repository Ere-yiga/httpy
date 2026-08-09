package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct{
	Title string `json:"title"`
	Author string `json:"author"`
	Pages uint `json:"pages"`
}

func BookHandler(w http.ResponseWriter, r *http.Request){

	b := Book {Title: "The beautitudes", Author: "---", Pages: 12}
	// I am enconding now
	json.NewEncoder(w).Encode(b)
}

func decodeBookHandler(w http.ResponseWriter, r *http.Request){
	var b Book
	json.NewDecoder(r.Body).Decode(&b)
	fmt.Println("Received", b.Author, b.Title, b.Pages)

	w.Write([]byte("Received book: " + b.Title))
}

func main(){
	fmt.Println("waiting...")
	http.HandleFunc("/Book", BookHandler)	
	http.HandleFunc("/decode", decodeBookHandler)
	http.ListenAndServe(":8080", nil)
}