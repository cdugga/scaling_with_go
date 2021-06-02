package controller

import (
	"encoding/json"
	"github.com/cdugga/bookmark/model"
	"github.com/cdugga/bookmark/service"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	locSVC service.LocationService = service.NewLocService()
)

func GetOrgById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	place := params["id"]

	var googlebook model.GoogleBook
	var location []byte
	var err error

	if location, err = locSVC.GetLocationById(place); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err:= json.Unmarshal(location, &googlebook); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, googlebook)

}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}