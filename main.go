package main

import (
	olxrealestatescrapper "github.com/Yom3n/webscrapper/olxRealEstateScrapper"
)

// Webscrap realestates from olx. Output it to csv file where you list basic data about the realestate
func main() {
	// c := colly.NewCollector()
	// // Find and visint all links
	// c.OnHTML("a[href]", func(h *colly.HTMLElement) {
	// 	h.Request.Visit(h.Attr("href"))
	// })

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	// c.Visit("http://go-colly.org/")
	// scrapper := olxrealestatescrapper.NewScrapper()
	// html, err := scrapper.GetPageHTMLContent("https://www.olx.pl/nieruchomosci/mieszkania/")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(html)
	olxrealestatescrapper.ScrapRealEstates()
}
