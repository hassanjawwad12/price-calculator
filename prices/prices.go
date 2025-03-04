package prices

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hassanjawwad12/price-calculator/conversion"
)

type TaxIncludedPriceJob struct {
	InputPrices       []float64
	TaxRate           float64
	TaxIncludedPrices map[string]float64
}

// Constructor Function
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Error opening file")
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	var lines []string

	// read each line of the file
	for scanner.Scan() {
		// append in the slice
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file")
		fmt.Println(err)
		file.Close()
		return
	}

	prices, err := conversion.StringstoFloat(lines)
	if err != nil {
		fmt.Println("Error converting string to float")
		fmt.Println(err)
		file.Close()
		return
	}

	job.InputPrices = prices
	file.Close()
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

	fmt.Println(result)
}
