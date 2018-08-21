package structure

import (
	"encoding/json"
	"sort"
	"fmt"
	"time"
)

func CalculateDaysUntilResetString(ballCount int) string {
	ballCount, days, seconds := CalculateDaysUntilReset(ballCount)
	// We have to calculate millis ourselves https://github.com/golang/go/issues/5491
	millis := int(seconds * 1e3)
	// https://golang.org/pkg/fmt/
	resultString := fmt.Sprintf("%d balls cycle after %d days.\nCompleted in %d milliseconds (%.3f seconds)\n", ballCount, days, millis, seconds)
	return resultString
}

func CalculateDaysUntilReset(ballCount int) (int, int, float64) {
	clock := NewBallClock(ballCount)
	days := 0
	start := time.Now()
	for {
		clock.TickDay()
		days++
		if clock.IsInitialOrdering() {
			break
		}
	}
	duration := time.Since(start)
	return clock.count, days, duration.Seconds()
}

type BallClock struct {
	Min     []int `json:"Min"`
	FiveMin []int `json:"FiveMin"`
	Hour    []int `json:"Hour"`
	Main    []int `json:"Main"`
	count   int
}

func NewBallClock(ballCount int) BallClock {
	ballClock := BallClock{
		count: ballCount,
	}
	ballClock.Reset()
	return ballClock
}

func (clock *BallClock) Reset() {
	clock.Min = []int{}
	clock.FiveMin = []int{}
	clock.Hour = []int{}
	clock.Main = []int{}
	for i := 1; i <= clock.count; i++ {
		clock.Main = append(clock.Main, i)
	}
}

func (clock *BallClock) JsonString() string {
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

func (clock *BallClock) TickHours(hours int) {
	for i := 0; i < hours; i++ {
		clock.TickHour()
	}
}

func (clock *BallClock) TickHour() {
	clock.TickMinutes(60)
}

func (clock *BallClock) TickMinutes(minutes int) {
	for i := 0; i < minutes; i++ {
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
