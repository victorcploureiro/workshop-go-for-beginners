package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ParseCsvFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("error while opening csv file: %w", err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading csv file: %w", err)
	}

	return records, nil
}

func main() {
	records, err := ParseCsvFile("testdata/users_001.csv")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(records)
}