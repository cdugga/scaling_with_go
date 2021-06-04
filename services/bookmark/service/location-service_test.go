package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cdugga/bookmark/mocks"
	"github.com/cdugga/bookmark/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func init() {
	HttpClient = &mocks.MockClient{} // override HttpClient instance
}

func TestRequestSuccess(t *testing.T){

	// make sure to reset original func once completed
	savedEnv := Env
	defer func() { Env = savedEnv}()
	Env = &mocks.MockEnv{}

	//  returns supplied val
	mocks.GetFunc = func(key string) interface{} {
		return "https://www.googleapis.com/books/v1/"
	}

	jsonStr := `{"kind":"books#volumes","totalItems":57,"items":[{"kind":"books#volume","id":"KUgKEAAAQBAJ","etag":"a553Qr5X5P4","selfLink":"https://www.googleapis.com/books/v1/volumes/KUgKEAAAQBAJ","volumeInfo":{"title":"Fortune'sDaughter(TheRockwoodChronicles,Book1)","subtitle":"","authors":["DillyCourt"],"publisher":"HarperCollins","publishedDate":"2021-06-10","description":"Don’tmissthebrand-newsix-partseriesfromtheNo.1SundayTimesbestsellingauthorDillyCourt!","industryIdentifiers":[{"type":"ISBN_13","identifier":"9780008435509"},{"type":"ISBN_10","identifier":"0008435502"}],"pageCount":528,"printType":"BOOK","categories":["Fiction"],"maturityRating":"NOT_MATURE","allowAnonLogging":false,"contentVersion":"1.1.1.0.preview.2","language":"un","previewLink":"http://books.google.ie/books?id=KUgKEAAAQBAJ\\u0026printsec=frontcover\\u0026dq=book+intitle:rockwood\\u0026hl=\\u0026cd=1\\u0026source=gbs_api","infoLink":"http://books.google.ie/books?id=KUgKEAAAQBAJ\\u0026dq=book+intitle:rockwood\\u0026hl=\\u0026source=gbs_api","canonicalVolumeLink":"https://books.google.com/books/about/Fortune_s_Daughter_The_Rockwood_Chronicl.html?hl=\\u0026id=KUgKEAAAQBAJ"},"saleInfo":{"country":"IE","saleability":"NOT_FOR_SALE","isEbook":false,"listPrice":{"amount":null,"currencyCode":""},"retailPrice":{"amount":null,"currencyCode":""},"buyLink":"","offers":null},"accessInfo":{"country":"IE","viewability":"PARTIAL","textToSpeechPermission":"ALLOWED_FOR_ACCESSIBILITY","epub":{"isAvailable":true},"webReaderLink":"http://play.google.com/books/reader?id=KUgKEAAAQBAJ\\u0026hl=\\u0026printsec=frontcover\\u0026source=gbs_api"},"searchInfo":{}}]}`
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(jsonStr)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	location := NewLocService()

	resp, err := location.GetLocationById("rockwood",1)
	var googlebook model.GoogleBook
	json.Unmarshal(resp, &googlebook)

	r, _ := json.Marshal(googlebook)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, jsonStr, string(r))
}


func TestNewRequestCreationFailure(t *testing.T) {

	// make sure to reset original func once completed
	savedEnv := Env
	defer func() { Env = savedEnv}()
	Env = &mocks.MockEnv{}

	//  returns supplied val
	mocks.GetFunc = func(key string) interface{} {
		return nil
	}

	jsonStr := `{}`
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(jsonStr)))
		return &http.Response{
			StatusCode: 500,
			Body:       r,
		}, nil
	}

	location := NewLocService()
	_, err := location.GetLocationById("someLoc", 1)

	assert.NotNil(t, err)
}

func TestHttpClientDoFailure(t *testing.T) {
	// make sure to reset original func once completed
	savedEnv := Env
	defer func() { Env = savedEnv}()
	Env = &mocks.MockEnv{}

	//  returns supplied val
	mocks.GetFunc = func(key string) interface{} {
		return "https://www.googleapis.com/books/v1/"
	}
	jsonStr := `{}`
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(jsonStr)))
		return &http.Response{
			StatusCode: 500,
			Body:       r,
		}, fmt.Errorf("HttpClient request failed")
	}

	location := NewLocService()
	_, err := location.GetLocationById("rockwood", 1)

	assert.NotNil(t, err)

}

func TestUTFEncoding(t *testing.T) {

	// make sure to reset original func once completed
	savedEnv := Env
	defer func() { Env = savedEnv}()
	Env = &mocks.MockEnv{}

	//  returns supplied val
	mocks.GetFunc = func(key string) interface{} {
		return "https://www.googleapis.com/books/v1/"
	}
	jsonStr := `λ**%&^$^%()&)(*&_)*)_*(*^^&£^$%^$*&&)*_)*(*_)(*`
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(jsonStr)))
		return &http.Response{
			StatusCode: 500,
			Body: r,
		}, nil
	}

	location := NewLocService()
	_, err := location.GetLocationById("rockwood", 1)

	assert.Nil(t, err)
}

func BenchmarkRequestSuccess(b *testing.B) {
	// make sure to reset original func once completed
	savedEnv := Env
	defer func() { Env = savedEnv}()
	Env = &mocks.MockEnv{}

	//  returns supplied val
	mocks.GetFunc = func(key string) interface{} {
		return "https://www.googleapis.com/books/v1/"
	}

	jsonStr := `{"kind":"books#volumes","totalItems":57,"items":[{"kind":"books#volume","id":"KUgKEAAAQBAJ","etag":"hY8HHuj3o34","selfLink":"https://www.googleapis.com/books/v1/volumes/KUgKEAAAQBAJ","volumeInfo":{"title":"Fortune's Daughter (The Rockwood Chronicles, Book 1)","subtitle":"","authors":["Dilly Court"],"publisher":"HarperCollins","publishedDate":"2021-06-10","description":"Don’t miss the brand-new six-part series from the No.1 Sunday Times bestselling author Dilly Court!","industryIdentifiers":[{"type":"ISBN_13","identifier":"9780008435509"},{"type":"ISBN_10","identifier":"0008435502"}],"readingModes":{"text":true,"image":false},"pageCount":528,"printType":"BOOK","categories":["Fiction"],"maturityRating":"NOT_MATURE","allowAnonLogging":false,"contentVersion":"1.1.1.0.preview.2","panelizationSummary":{"containsEpubBubbles":false,"containsImageBubbles":false},"imageLinks":{"smallThumbnail":"http://books.google.com/books/content?id=KUgKEAAAQBAJ\u0026printsec=frontcover\u0026img=1\u0026zoom=5\u0026edge=curl\u0026source=gbs_api","thumbnail":"http://books.google.com/books/content?id=KUgKEAAAQBAJ\u0026printsec=frontcover\u0026img=1\u0026zoom=1\u0026edge=curl\u0026source=gbs_api"},"language":"un","previewLink":"http://books.google.ie/books?id=KUgKEAAAQBAJ\u0026printsec=frontcover\u0026dq=book+intitle:rockwood\u0026hl=\u0026cd=1\u0026source=gbs_api","infoLink":"http://books.google.ie/books?id=KUgKEAAAQBAJ\u0026dq=book+intitle:rockwood\u0026hl=\u0026source=gbs_api","canonicalVolumeLink":"https://books.google.com/books/about/Fortune_s_Daughter_The_Rockwood_Chronicl.html?hl=\u0026id=KUgKEAAAQBAJ"},"saleInfo":{"country":"IE","saleability":"NOT_FOR_SALE","isEbook":false,"listPrice":{"amount":null,"currencyCode":""},"retailPrice":{"amount":null,"currencyCode":""},"buyLink":"","offers":null},"accessInfo":{"country":"IE","viewability":"PARTIAL","embeddable":true,"publicDomain":false,"textToSpeechPermission":"ALLOWED_FOR_ACCESSIBILITY","epub":{"isAvailable":true},"pdf":{"isAvailable":false},"webReaderLink":"http://play.google.com/books/reader?id=KUgKEAAAQBAJ\u0026hl=\u0026printsec=frontcover\u0026source=gbs_api","accessViewStatus":"SAMPLE","quoteSharingAllowed":false},"searchInfo":{"textSnippet":"Don’t miss the brand-new six-part series from the No.1 Sunday Times bestselling author Dilly Court!"}}]}`
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(jsonStr)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	location := NewLocService()


	for i := 0; i < b.N; i++ {
		resp, _ := location.GetLocationById("rockwood", 1)

		var googlebook model.GoogleBook
		json.Unmarshal(resp, &googlebook)
		//
		//json.Marshal(googlebook)

	}
}
