package main

import "fmt"
import "os"

func main() {
	file, err := os.ReadFile("testdata/users_001.csv")
	if err != nil {
		return
	}

	fmt.Println(string(file))
}