package structure

import (
	"encoding/json"
	"sort"
)

type BallClock struct {
	Min     []int `json:"Min"`
	FiveMin []int `json:"FiveMin"`
	Hour    []int `json:"Hour"`
	Main    []int `json:"Main"`
}

func NewBallClock(ballCount int) BallClock {
	ballClock := BallClock{
		Min:     []int{},
		FiveMin: []int{},
		Hour:    []int{},
		Main:    []int{},
	}
	for i := 1; i <= ballCount; i++ {
		ballClock.Main = append(ballClock.Main, i)
	}
	return ballClock
}

func (clock *BallClock) String() string {
	bytes, _ := json.Marshal(clock)
	return string(bytes)
}

func (clock *BallClock) IsInitialOrdering() bool {
	return sort.IntsAreSorted(clock.Main)
}

func (clock *BallClock) TickDay() {
	for i := 0; i < 24; i++ {
		clock.TickHour()
	}
}

func (clock *BallClock) TickHour() {
	for i := 0; i < 60; i++ {
		clock.TickMinute()
	}
}

func (clock *BallClock) TickMinute() {
	var nextBall int
	nextBall, clock.Main = clock.Main[0], clock.Main[1:]
	clock.Min = append(clock.Min, nextBall)

	var ballCarryOver int
	ballCarryOver, clock.Min, clock.Main = checkBallCarryOver(5, clock.Min, clock.Main)
	if ballCarryOver >= 1 {
		clock.FiveMin = append(clock.FiveMin, ballCarryOver)
		ballCarryOver, clock.FiveMin, clock.Main = checkBallCarryOver(12, clock.FiveMin, clock.Main)
	}
	if ballCarryOver >= 1 {
		clock.Hour = append(clock.Hour, ballCarryOver)
		ballCarryOver, clock.Hour, clock.Main = checkBallCarryOver(12, clock.Hour, clock.Main)
	}
	if ballCarryOver >= 1 {
		clock.Main = append(clock.Main, ballCarryOver)
	}
}

func checkBallCarryOver(threshold int, currentBuffer []int, mainBuffer []int) (int, []int, []int) {
	ballCarryOver := -1
	length := len(currentBuffer)
	if length == threshold {
		ballCarryOver = currentBuffer[length-1]
		for i := length - 2; i >= 0; i-- {
			mainBuffer = append(mainBuffer, currentBuffer[i])
		}
		currentBuffer = []int{}
	}
	return ballCarryOver, currentBuffer, mainBuffer
}
