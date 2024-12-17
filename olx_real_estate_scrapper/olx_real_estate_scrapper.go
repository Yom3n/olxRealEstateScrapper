package olxrealestatescrapper

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Yom3n/webscrapper/models"
	"github.com/Yom3n/webscrapper/web_scrapper"
)

type OlxRealEstateScrapper struct {
	WebScrapper *web_scrapper.WebScrapper
}

func (o *OlxRealEstateScrapper) ScrapRealEstates() {
	url := getUrlWithPage(1)
	firstPage, maxPages := o.scrapSinglePage(url, nil)
	firstPage.Print()
	if maxPages == -1 {
		log.Fatal("Coudln't scap num pages")
	}
	if maxPages == 1 {
		// TODO Return first page value
		return
	}

	channel := make(chan models.RealEstatesRecrods)
	for page := 2; page <= maxPages; page++ {
		go o.scrapSinglePage(getUrlWithPage(page), channel)
	}

	for page := 2; page <= maxPages; page++ {
		estates := <-channel
		estates.Print()
	}
}

func getUrlWithPage(page int) string {
	pageStr := strconv.Itoa(page)
	return "https://www.olx.pl/nieruchomosci/mieszkania/sprzedaz/?page=" + pageStr
}

// Provide full url with params, pages etc.
// ex: https://www.olx.pl/nieruchomosci/mieszkania/sprzedaz/?page=1
func (o *OlxRealEstateScrapper) scrapSinglePage(url string, channel chan models.RealEstatesRecrods) (estatesPage models.RealEstatesRecrods, numPages int) {
	html, err := o.WebScrapper.GetPageHTMLContent(url)
	if err != nil {
		log.Fatal()
		return
	}
	res := getRealEstatesFromHtml(html)
	_, maxPageStr := getValueFromHtml(html, numPaginationPagesKey, numPaginationPagesEndKey)
	numPaginationPages, err := strconv.Atoi(maxPageStr)
	if err != nil {
		fmt.Println(err)
		if channel != nil {
			channel <- res
		}
		return res, -1
	}
	if channel != nil {
		channel <- res
	}
	return res, numPaginationPages
}

// After this key starts ad title. Ends with "
const titleKey = `\"title\":\"`
const titleEndKey = `\",\"`

// /\"regularPrice\":{\"value\":349000,
const priceKey = `\"regularPrice\":{\"value\":`
const priceEndKey = ","

const areaKey = `\"Powierzchnia\",\"type\":\"input\",\"value\":\"`
const areaEndKey = ` m²\",`

const numPaginationPagesKey = `...</li><li data-testid="pagination-list-item" aria-label="Page `
const numPaginationPagesEndKey = `"`

func getRealEstatesFromHtml(html string) models.RealEstatesRecrods {
	realEstates := models.RealEstatesRecrods{}
	for {
		var title string
		var price string
		var area float32
		html, title = getValueFromHtml(html, titleKey, titleEndKey)
		if title == "" {
			break
		}
		html, price = getValueFromHtml(html, priceKey, priceEndKey)
		if price == "" {
			continue
		}

		var areaStr = ""
		html, areaStr = getValueFromHtml(html, areaKey, areaEndKey)
		if areaStr == "" {
			continue
		}
		areaStr = strings.Replace(areaStr, ",", ".", 1)
		parsedArea, err := strconv.ParseFloat(areaStr, 32)
		if err != nil {
			fmt.Println(err)
			continue
		}
		area = float32(parsedArea)

		priceInt, err := strconv.Atoi(price)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		realEstates = append(realEstates, models.RealEstate{
			Title:      title,
			PriceZloty: priceInt,
			AreaInM2:   area,
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
