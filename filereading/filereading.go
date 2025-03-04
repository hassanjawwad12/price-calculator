package filereading

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("error opening file")
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
		return nil, errors.New("error reading file")
	}
	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to convert data to JSON")
	}

	file.Close()
	return nil
}
