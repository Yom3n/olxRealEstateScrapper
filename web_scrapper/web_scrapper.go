package web_scrapper

import (
	"io"
	"net/http"
)

type WebScrapper struct {
}

func NewWebScrapper() WebScrapper {
	return WebScrapper{}
}

func (s *WebScrapper) GetPageHTMLContent(url string) (content string, err error) {
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
