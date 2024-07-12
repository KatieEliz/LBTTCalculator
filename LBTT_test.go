package main

import (
	"testing"
)

func TestWhenPriceIsUnderThreashold(t *testing.T) {
	price := 100000
	expectedTax := 0.0

	result := CalculateTaxRate(price)

	if result != expectedTax {
		t.Error("Unexpected tax rate")
	}

}
func TestFirstThreashold(t *testing.T) {
	price := 145000
	expectedTax := 0.0

	result := CalculateTaxRate(price)

	if result != expectedTax {
		t.Error("Unexpected tax rate")
	}
}
func TestBetweenFirstAndSecondThreashold(t *testing.T) {
	price := 165000

	result := CalculateTaxRate(price)
	expectedTax := 400

	if result != float64(expectedTax) {
		t.Error("Unexpected tax rate")
	}
}
func TestBetweenSecondAndThirdThreashold(t *testing.T) {
	price := 300000

	result := CalculateTaxRate(price)
	expectedTax := 4600

	if result != float64(expectedTax) {
		t.Error("Unexpected tax rate")
	}
}
func TestBetweenThirdAndFourthThreashold(t *testing.T) {
	price := 400000

	result := CalculateTaxRate(price)
	expectedTax := 13350

	if result != float64(expectedTax) {
		t.Error("Unexpected tax rate")
	}
}
func TestLastThreashold(t *testing.T) {
	price := 750000

	result := CalculateTaxRate(price)
	expectedTax := 48350

	if result != float64(expectedTax) {
		t.Error("Unexpected tax rate")
	}
}

/*func TestOutsideThreshold(t *testing.T) {
	price := 800000

	result := CalculateTaxRate(price)
	expectedRate := 60000.0

	if result != expectedRate {
		t.Error("Unexpected tax")
	}
}*/
