package main

import "fmt"

type playerID string

func newPlayerID(id string) (playerID, error) {
	// playerIDのバリデーションはplayer登録時に行うべきなので、こちらでは省略
	// あくまでrankingの計算をする
	return playerID(id), nil
}

type playLog struct {
	timeStamp string
	score     int
	playerID  playerID
}

func newPlayLog(score int, timestamp, playerID string) (*playLog, error) {
	if score < 0 {
		return nil, fmt.Errorf("score must be more than 0")
	}

	pID, err := newPlayerID(playerID)
	if err != nil {
		return nil, err
	}

	playLog := &playLog{
		timeStamp: timestamp,
		score:     score,
		playerID:  pID,
	}
	return playLog, nil
}

type playerIDToAverage map[playerID]int

type averageToPlayerIDs map[int][]playerID

func (a averageToPlayerIDs) getKeys() []int {
	keys := make([]int, 0, len(a))
	for k := range a {
		keys = append(keys, k)
	}
	return keys
}
