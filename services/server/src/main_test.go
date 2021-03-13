package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello from the other side/"
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