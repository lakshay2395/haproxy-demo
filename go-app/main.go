package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"message": "Welcome to Dockerized app",
		}
		json.NewEncoder(rw).Encode(response)
	})

	router.HandleFunc("/bye", func(rw http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"message": "Bye bye",
		}
		json.NewEncoder(rw).Encode(response)
	})

	log.Println("Server is running!")
	fmt.Println(http.ListenAndServe(":8081", router))
}
