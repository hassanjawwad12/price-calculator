package conversion

import (
	"errors"
	"strconv"
)

func StringstoFloat(input []string) ([]float64, error) {

	floats := make([]float64, len(input))

	for lineIndex, line := range input {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			return nil, errors.New("error converting string to float")
		}

		floats[lineIndex] = floatPrice
		// floats = append(floats, floatPrice)
	}
	return floats, nil
}
