package olxrealestatescrapper

import (
	"fmt"
	"log"
	"strings"
)

type RealEstate struct {
	title      string
	priceZloty int
	areaInM2   int
}

func ScrapRealEstates() {
	scrapper := NewScrapper()
	html, err := scrapper.GetPageHTMLContent("https://www.olx.pl/nieruchomosci/mieszkania/")
	if err != nil {
		log.Fatal()
		return
	}
	getRealEstatesFromHtml(html)

}

// After this key starts ad title. Ends with "
const titleKey = `\"title\":\"`

func getRealEstatesFromHtml(html string) []RealEstate {
	realEstates := []RealEstate{}
	for {
		titleIndex := strings.Index(html, titleKey)
		if titleIndex == -1 {
			return realEstates
		}
		html = html[titleIndex+len(titleKey):]
		endIndex := strings.Index(html, "\"") - 1
		if endIndex == -1 {
			log.Fatal("Missing closing \" for title")
		}
		title := html[:endIndex]
		fmt.Println(title)
		html = html[endIndex+1:]
	}

}
