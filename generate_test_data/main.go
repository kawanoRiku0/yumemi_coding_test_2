package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 1 {
		log.Fatal("please input a number of test data rows ")
	}
	numRows, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}
	generateTestCSV(numRows)
}

func generateTestCSV(n int) {
	var rows string

	for i := 0; i < n; i++ {
		// 今回の課題には関係ないので、定数
		timeStamp := "2021/01/01 12:00"
		playerID := NewRandomPlayerID()
		score := rand.Intn(30000)
		row := fmt.Sprintf("%s,%s,%d\n", timeStamp, playerID, score)
		rows += row
	}

	file, err := os.Create("../yumemi_test/data/ranking.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(rows)
}

func NewRandomPlayerID() string {
	playerID := "player"

	numOfDigits := 4
	for i := 0; i < numOfDigits; i++ {
		playerID += strconv.Itoa(rand.Intn(10))
	}

	return playerID
}
