package main

import (
	"github.com/hassanjawwad12/price-calculator/prices"
)

func main() {

	taxRates := []float64{0.1, 0.2, 0.3, 0.4, 0.5}

	for _, taxRate := range taxRates {
		pricejob := prices.NewTaxIncludedPriceJob(taxRate)
		// map created for each tax rate
		pricejob.Process()
	}
}
