package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	readText := string(data)
	readFloat, err := strconv.ParseFloat(readText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return readFloat, nil
}

func WriteFloatToFile(data float64, fileName string) {
	dataText := fmt.Sprint(data)
	os.WriteFile(fileName, []byte(dataText), 0644)
}
