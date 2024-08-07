package main

import (
	"fmt"
)

func main() {
	records, err := ParseCsvFile("testdata/users_001.csv")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(records)
}