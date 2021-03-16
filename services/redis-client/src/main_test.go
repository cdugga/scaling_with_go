package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/alicebob/miniredis/v2"
)

func TestHandler(t *testing.T) {

	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	s.Set("name", "colin")

	defer s.Close()

	path := fmt.Sprintf("/value/%s", "name")

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/value/{key}", FetchKeyHandler)
	router.ServeHTTP(rr, req)


	if rr.Code != http.StatusOK {
		t.Errorf("handler should have failed : got %v want %v",
			rr.Code, http.StatusOK)
	}

	expected := "Requested value colin"
	if rr.Body.String() != expected {
		t.Errorf("handler retruned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

}



//func TestServer(t *testing.T){
//	var resp string
//	var apiStub = httptest.NewServer(http.HandlerFunc(
//		func(w http.ResponseWriter, r *http.Request) {
//			switch r.RequestURI {
//			case "/" :
//				resp = string(http.StatusOK)
//			case "/other" :
//				resp = string(http.StateClosed)
//			default:
//				http.Error(w, "not found", http.StatusNotFound)
//			}
//			w.Write([]byte(resp))
//		}))
//
//	defer apiStub.Close()
//
//
//}