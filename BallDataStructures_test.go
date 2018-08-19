package main

import (
	"testing"
	"github.com/paul-nelson-baker/ball-clock-simulator/structure"
)

func TestClockMatchesPDFIterations(t *testing.T) {
	ballClock := structure.NewBallClock(30)
	for i := 0; i < 325; i++ {
		ballClock.TickMinute()
	}
	testSlicesEqual(t, ballClock.Min, []int{})
	testSlicesEqual(t, ballClock.FiveMin, []int{22, 13, 25, 3, 7})
	testSlicesEqual(t, ballClock.Hour, []int{6, 12, 17, 4, 15})
	testSlicesEqual(t, ballClock.Main, []int{11, 5, 26, 18, 2, 30, 19, 8, 24, 10, 29, 20, 16, 21, 28, 1, 23, 14, 27, 9})
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
	if days != 15 {
		t.FailNow()
	}
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
	if days != 378 {
		t.FailNow()
	}
}

func testSlicesEqual(t *testing.T, a, b []int) {
	if (a == nil) != (b == nil) {
		t.FailNow()
	}
	if len(a) != len(b) {
		t.FailNow()
	}
	for i := range a {
		if a[i] != b[i] {
			t.FailNow()
		}
	}
}
