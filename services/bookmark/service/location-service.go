package service

import (
	"fmt"
	chttp "github.com/cdugga/bookmark/http"
	"io"
	"net/http"
)

type LocationService interface {
	GetLocationById(place string) ([]byte, error)
}

var (
	HttpClient = chttp.Client
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

	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=book+intitle:%s&maxResults=1", place)

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