package models

import "fmt"

type RealEstatesRecrods []RealEstate

func (r *RealEstatesRecrods) Print() {
	for i, realEstate := range *r {
		fmt.Println(i+1, ".  ", realEstate.Title, "  ", realEstate.AreaInM2, "m2 ", realEstate.PriceZloty, "zł")
	}
}

type RealEstate struct {
	Title      string
	PriceZloty int
	AreaInM2   float32
}
