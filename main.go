package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func SelectWinners(records, [][]string, nWinners int) [][]string {
	rand.Shuffle(len(records), func (i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	return records[:nWinners]
}

func main() {
	userCsvFile := flag.String("file", "testdata/users_001.csv", "Path to users file")
	nWinners := flag.Int("winners", 6, "Number of winners to select")
	records, err := ParseCsvFile("testdata/users_001.csv")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	winners :=SelectWinners(records, 3)
	for i, winner := range winners {
		fmt.Println("Ganhador: ", i, " ", )
	}
}

// https://github.com/reneepc/workshop-go-for-beginners repositorio de como deveria ser