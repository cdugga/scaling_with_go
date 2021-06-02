package mocks

import "net/http"

var (
	// GetDoFunc fetches the mock client's `Do` func
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	// coming soon!
	return GetDoFunc(req)
}