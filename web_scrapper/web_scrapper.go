package web_scrapper

import (
	"fmt"
	"io"
	"net/http"
)

// / Create HttpClient interface for testing purpose
type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type WebScrapper struct {
	client HttpClient
}

func NewWebScrapper(client HttpClient) *WebScrapper {
	return &WebScrapper{client: client}
}

const maxReadSize = 10 * 1024 * 1024 // 10mb

func (s *WebScrapper) GetPageHTMLContent(url string) (content []byte, err error) {
	resp, err := s.client.Get(url)
	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("failed to fetch %s. Status code: %s", url, resp.Status)
	}
	if err != nil {
		return []byte{}, fmt.Errorf("failed to fetch %s: %w", url, err)
	}

	defer resp.Body.Close()
	limitedReader := io.LimitReader(resp.Body, maxReadSize)
	html, err := io.ReadAll(limitedReader)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read data from %s: %w", url, err)
	}

	return html, nil
}
