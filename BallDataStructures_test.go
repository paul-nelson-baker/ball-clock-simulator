package main

import (
	"testing"
	"github.com/paul-nelson-baker/ball-clock-simulator/structure"
	"fmt"
)

func TestClockMatchesPDFIterations(t *testing.T) {
	ballClock := structure.NewBallClock(30)
	for i := 0; i < 325; i++ {
		ballClock.TickMinute()
	}
	fmt.Println(ballClock.String())
}

func TestClockCycleDaysOne(t *testing.T) {
	ballClock := structure.NewBallClock(30)
	days := 0
	for {
		ballClock.TickDay()
		days++
		if ballClock.IsInitialOrdering() {
			break
		}
	}
	fmt.Printf("Days before reset: %d\n", days)
	fmt.Println(ballClock.String())
}

func TestClockCycleDaysTwo(t *testing.T) {
	ballClock := structure.NewBallClock(45)
	days := 0
	for {
		ballClock.TickDay()
		days++
		if ballClock.IsInitialOrdering() {
			break
		}
	}
	fmt.Printf("Days before reset: %d\n", days)
	fmt.Println(ballClock.String())
}
