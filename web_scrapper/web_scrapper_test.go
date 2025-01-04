package web_scrapper

import (
	"net/http"
	"testing"
)

type HttpClientSuccess struct {
}

func (c *HttpClientSuccess) Get(url string) (resp *http.Response, err error) {
	return &http.Response{StatusCode: 200, Body: http.NoBody}, nil

}
func TestWebScrapperHappyPath(t *testing.T) {
	httpMock := &HttpClientSuccess{}
	sut := NewWebScrapper(httpMock)
	_, err := sut.GetPageHTMLContent("https://someUrl.com/")

	if err != nil {
		t.Fatal("Error should be nil")
	}

}
