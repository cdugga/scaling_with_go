package service

import (
	"context"
	"fmt"
	"github.com/cdugga/bookmark/env"
	chttp "github.com/cdugga/bookmark/http"
	"io"
	"net/http"
	"time"
)

type LocationService interface {
	GetLocationById(place string, maxResults int) ([]byte, error)
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

func (s *locService) GetLocationById(place string, maxResults int) ([]byte, error) {

	apiPath := Env.Get(API_PATH_KEY)
	url := fmt.Sprintf("%svolumes?q=book+intitle:%s&maxResults=%d", apiPath, place, maxResults)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("get %s failed : %v", url, err)
	}

	// set a per request timeout 15 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15)
	defer cancel()
	request = request.WithContext(ctx)

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