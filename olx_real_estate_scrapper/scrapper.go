package olxrealestatescrapper

import (
	"io"
	"net/http"
)

type Scrapper struct {
}

func NewScrapper() Scrapper {
	return Scrapper{}
}

func (s *Scrapper) GetPageHTMLContent(url string) (content string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	html, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(html), nil
}
