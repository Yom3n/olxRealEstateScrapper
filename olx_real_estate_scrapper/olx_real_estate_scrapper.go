package olxrealestatescrapper

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Yom3n/webscrapper/realEstateRecords"
)



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
const titleEndKey = "\""

// /\"regularPrice\":{\"value\":349000,
const priceKey = `\"regularPrice\":{\"value\":`
const priceEndKey = ","

func getRealEstatesFromHtml(html string) []realEstateRecords.RealEstate {
	realEstates := []realEstateRecords.RealEstate{}
	for {
		var title string
		var price string
		html, title = getValueFromHtml(html, titleKey, titleEndKey)
		if title == "" {
			break
		}
		html, price = getValueFromHtml(html, priceKey, priceEndKey)
		if price == "" {
			continue
		}
		fmt.Println(title)
		fmt.Println(price)
		priceInt, err := strconv.Atoi(price)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		realEstates = append(realEstates, realEstateRecords.RealEstate{
			Title:      title,
			PriceZloty: priceInt,
			AreaInM2:   0,
		})
	}
	return realEstates
}

// / Gets single value from the html code that contains json
// / key is the json key of a lookup value, with all characters that are before value. For example  `\"title\":\"`
// / end key is a character, or group of characters taht marks end of value. For example for \"regularPrice\":{"value": 35000, "someOtherKey":` the "," i end value
func getValueFromHtml(html string, key string, endKey string) (htmlOutput string, value string) {
	keyIndex := strings.Index(html, key)
	if keyIndex == -1 {
		// Key not found
		return html, ""
	}
	html = html[keyIndex+len(key):]
	endIndex := strings.Index(html, endKey)
	output := html[:endIndex]
	html = html[endIndex+1:]
	return html, output
}
