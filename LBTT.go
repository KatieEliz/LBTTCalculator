package main

var threasholds = []int{145000, 250000, 325000, 750000}
var rates = []float64{0, 0.02, 0.05, 0.10, 0.12}

func CalculateTaxRate(priceOfHouse int) float64 {
	var tax float64

	switch {
	case priceOfHouse <= threasholds[0]:
		tax = 0
	case priceOfHouse <= threasholds[1]:
		tax = float64(priceOfHouse-threasholds[0]) * rates[1]
	case priceOfHouse <= threasholds[2]:
		tax = float64(threasholds[1]-threasholds[0])*rates[1] + float64(priceOfHouse-threasholds[1])*rates[2]
	case priceOfHouse <= threasholds[3]:
		tax = float64(threasholds[1]-threasholds[0])*rates[1] + float64(threasholds[2]-threasholds[1])*rates[2] + float64(priceOfHouse-threasholds[2])*rates[3]
	default:
		tax = float64(threasholds[1]-threasholds[0])*rates[1] + float64(threasholds[2]-threasholds[1])*rates[2] + float64(threasholds[3]-threasholds[2])*rates[3] + float64(priceOfHouse-threasholds[3])*rates[4]
	}
	return tax
}
