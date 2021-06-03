package http

import (
	"net/http"
	"time"
)

var (
	Client HTTPClient
)

type HTTPClient interface {
	Do(req *http.Request)(*http.Response, error)
}

func init(){
	Client = &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Minute, // sets timeout for all requests, use context built by Google for more granular approach
	}
}