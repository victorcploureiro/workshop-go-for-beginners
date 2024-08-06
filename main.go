package main

import "fmt"
import "os"
import "encoding/csv"

func main() {
	file, err := os.Open("testdata/users_001.csv")
	if err != nil {
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	fmt.Println(records)
	file.Close()
}