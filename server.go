package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const PORT string = ":8000"
	router := mux.NewRouter()

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Up and running...")
	})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/post", addPost).Methods("POST")

	log.Println("Server listening on port: ", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
