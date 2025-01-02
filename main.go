package main

import (
	olxrealestatescrapper "github.com/Yom3n/webscrapper/olx_real_estate_scrapper"
	"github.com/Yom3n/webscrapper/web_scrapper"
	"net/http"
	"time"
)

// Webscrap realestates from olx. Output it to csv file where you list basic data about the realestate
func main() {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	scrapper := web_scrapper.NewWebScrapper(httpClient)
	olxScrapper := olxrealestatescrapper.NewOlxRealEstatesScrapper(
		scrapper,
	)
	realEstates := olxScrapper.ScrapRealEstates()
	realEstates.Print()
}
