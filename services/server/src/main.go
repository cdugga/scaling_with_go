package main

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {

	fmt.Println("Starting server")

	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello from the other side", r.URL.Path)
}
