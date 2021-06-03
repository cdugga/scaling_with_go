package service

import (
	"fmt"
	"github.com/cdugga/bookmark/env"
	chttp "github.com/cdugga/bookmark/http"
	"io"
	"net/http"
)

type LocationService interface {
	GetLocationById(place string) ([]byte, error)
}

const (
	API_PATH_KEY="GOOGLE_BOOKS_API"
)

var (
	HttpClient = chttp.Client
	Env env.Provider = env.NewEnv()
)

/*
 *	Org service layer to help interaction between org controller and databse.
**/
type locService struct {
}

func NewLocService() LocationService {
	return &locService{}
}

func (s *locService) GetLocationById(place string) ([]byte, error) {

	//apiPath := Env.Get(API_PATH_KEY)
	apiPath := "https://www.googleapis.com/books/v1/"
	url := fmt.Sprintf("%svolumes?q=book+intitle:%s&maxResults=1", apiPath, place)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("get %s failed : %v", url, err)
	}

	resp, err := HttpClient.Do(request)
	if err != nil{
		return nil, fmt.Errorf("httpClient.Do %s as response body : %v", resp.Body, err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil{
		return nil, fmt.Errorf("readAll %s as response body : %v", resp.Body, err)
	}

	return body, nil
}