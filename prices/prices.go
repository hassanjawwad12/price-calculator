package prices

import (
	"fmt"

	"github.com/hassanjawwad12/price-calculator/conversion"
	"github.com/hassanjawwad12/price-calculator/filereading"
)

type TaxIncludedPriceJob struct {
	InputPrices       []float64
	TaxRate           float64
	TaxIncludedPrices map[string]string
}

// Constructor Function
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := filereading.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringstoFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {

	// Load data from file
	job.LoadData()

	result := make(map[string]string)
	// key is price and value is tax applied price
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	filereading.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
}
