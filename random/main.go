package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	now := time.Now()
	baseMinutes := []int { 400, 420, 440, 460, 480, 500, 520, 540, 560, 580, 600, 620, 640, 660, 680, 700 }

	for i := 0; i < 20; i++ {
		date := now.AddDate(0, 0, (i + 1))
		for j := 0; j < 5; j++ {
			baseMinutesIndex := rand.Intn(len(baseMinutes))
			adjust := rand.Intn(100)
			minutes := baseMinutes[baseMinutesIndex] + adjust
			fmt.Printf("(%d, '%s', %d), ", j+1, date.Format("2006-01-02"), minutes)
		}
		fmt.Println("")
	}
}
