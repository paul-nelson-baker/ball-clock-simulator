package main

import (
	"testing"
	"github.com/paul-nelson-baker/ball-clock-simulator/structure"
	"fmt"
)

func TestClockMatchesPDFIterations(t *testing.T) {
	ballClock := structure.NewBallClock(30)
	ballClock.TickMinutes(325)
	fmt.Println(ballClock.JsonString())
	testSlicesEqual(t, ballClock.Min, []int{})
	testSlicesEqual(t, ballClock.FiveMin, []int{22, 13, 25, 3, 7})
	testSlicesEqual(t, ballClock.Hour, []int{6, 12, 17, 4, 15})
	testSlicesEqual(t, ballClock.Main, []int{11, 5, 26, 18, 2, 30, 19, 8, 24, 10, 29, 20, 16, 21, 28, 1, 23, 14, 27, 9})
}

func TestClockCycleDaysOne(t *testing.T) {
	_, daysUntilReset, _ := structure.CalculateDaysUntilReset(30)
	if daysUntilReset != 15 {
		t.FailNow()
	}
}

func TestClockCycleDaysTwo(t *testing.T) {
	_, daysUntilReset, _ := structure.CalculateDaysUntilReset(45)
	if daysUntilReset != 378 {
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
