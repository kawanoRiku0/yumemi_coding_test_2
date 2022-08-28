package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 1 {
		log.Fatal("please input a file name")
	}
	fileName := args[1]

	playLogs := generatePlayLogsFromCSVFile(fileName)
	playerIDToAverage := generatePlayerIDToAverageMap(playLogs)
	averageToPlayerIDs := generateAverageToPlayerIDsMap(playerIDToAverage)
	printRanking(averageToPlayerIDs)

}

func generatePlayerIDToAverageMap(playLogs []*playLog) playerIDToAverage {
	// 平均点を計算するために、playerIDの出現回数を計算
	playerIDCounter := make(map[playerID]int)

	playerIDToAverage := make(playerIDToAverage)
	for _, playLog := range playLogs {
		playerIDCounter[playLog.playerID]++
		playerIDToAverage[playLog.playerID] += playLog.score
	}

	// 平均点を計算して再代入
	for k, v := range playerIDToAverage {
		ave := float64(v) / float64(playerIDCounter[k])
		roundedAve := math.Round(ave)
		playerIDToAverage[k] = int(roundedAve)
	}

	return playerIDToAverage
}

func generatePlayLogsFromCSVFile(fileName string) []*playLog {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	playLogs := make([]*playLog, 0, len(rows))
	for i, row := range rows {
		// 一行目はheaderなのでスキップ
		if i == 0 {
			continue
		}

		score, err := strconv.Atoi(row[2])
		if err != nil {
			log.Fatal(err)
		}

		playLog, err := newPlayLog(score, row[0], row[1])
		if err != nil {
			fmt.Printf("error index %d : %s\n", i, err)
		}
		playLogs = append(playLogs, playLog)
	}
	return playLogs
}

func generateAverageToPlayerIDsMap(playerIDToAverage playerIDToAverage) averageToPlayerIDs {
	averageToPlayerIDs := make(averageToPlayerIDs)
	for ID, ave := range playerIDToAverage {
		averageToPlayerIDs[ave] = append(averageToPlayerIDs[ave], ID)
	}
	return averageToPlayerIDs
}

func printRanking(averageToPlayerIDs averageToPlayerIDs) {
	limit := 10
	printCount := 0
	rankCount := 1

	fmt.Println("rank, player_id, mean_score")

	scores := averageToPlayerIDs.getKeys()
	// 降順にソート
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))

	for _, score := range scores {
		if printCount >= limit {
			break
		}

		playerIDs := averageToPlayerIDs[score]

		for _, ID := range playerIDs {
			fmt.Printf("%d, %s, %d\n", rankCount, ID, score)
			printCount++
		}
		rankCount++
	}
}
