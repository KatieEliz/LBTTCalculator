package main

type taxBand struct {
	lower int
	upper int
	rate  float64
}

var bands []taxBand = []taxBand{
	{lower: 0, upper: 145000, rate: 0.0},
	{lower: 145000, upper: 250000, rate: 0.02},
	{lower: 250000, upper: 325000, rate: 0.05},
	{lower: 325000, upper: 750000, rate: 0.10},
	{lower: 750000, upper: 9999999, rate: 0.12},
}

func (band taxBand) Includes(price int) bool {
	return price >= band.lower && price <= band.upper
}

func (band taxBand) isAbove(price int) bool {
	return price > band.upper
}

func (band taxBand) taxWithinBand(price int) float64 {
	return float64(price-band.lower) * band.rate
}

func (band taxBand) maxTaxForBand(price int) float64 {
	return float64(band.upper-band.lower) * band.rate
}

func (band taxBand) taxDue(price int) float64 {
	if band.Includes(price) {
		return band.taxWithinBand(price)
	}

	if band.isAbove(price) {
		return band.maxTaxForBand(price)
	}
	return 0
}

func CalculateTaxRate(priceOfHouse int) float64 {
	tax := 0.0

	for _, band := range bands {
		tax += band.taxDue(priceOfHouse)
	}
	return tax

}
