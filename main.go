package main

import (
	olxrealestatescrapper "github.com/Yom3n/webscrapper/olx_real_estate_scrapper"
	"github.com/Yom3n/webscrapper/web_scrapper"
)

// Webscrap realestates from olx. Output it to csv file where you list basic data about the realestate
func main() {
	scrapper := web_scrapper.NewWebScrapper()
	olxScrapper := olxrealestatescrapper.OlxRealEstateScrapper{
		WebScrapper: &scrapper,
	}
	olxScrapper.ScrapRealEstates()
}
