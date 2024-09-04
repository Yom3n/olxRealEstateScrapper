package realEstateRecords

import "fmt"

type RealEstatesRecrods []RealEstate

func (r *RealEstatesRecrods) print() {
	for i, realEstate := range *r {
		fmt.Println(i+1, ".  ", realEstate.Title, "  ", realEstate.PriceZloty, "zł")
	}
}

type RealEstate struct {
	Title      string
	PriceZloty int
	AreaInM2   int
}