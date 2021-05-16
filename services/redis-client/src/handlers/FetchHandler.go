package handlers

import (
	"fmt"
	"github.com/cdugga/scaling_with_go/redisclient/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func FetchKeyHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	key, err := vars["key"]

	if err != true {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	log.Printf("Fetching key %s" , key)

	data := data.Datasource

	val, err1 := data.Get(key)
	if err1 != nil {
		panic(err1)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Requested value ", val)
}