package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Human struct {
	Name string  `json:"name"`
	Age  float64 `json:"age"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func humanHandler(w http.ResponseWriter, r *http.Request) {

	// Encoding
	p := Human{Name: "Shalom", Age: 49}
	json.NewEncoder(w).Encode(p)
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	
	// Decoding
	var p Human
	json.NewDecoder(r.Body).Decode(&p)
	fmt.Println("Received:", p.Name, p.Age)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/human", humanHandler)
	http.HandleFunc("/decode", decodeHandler)

	fmt.Println("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}