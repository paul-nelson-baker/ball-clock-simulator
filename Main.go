package main

import (
	"github.com/paul-nelson-baker/ball-clock-simulator/structure"
	"fmt"
)

func main() {
	ballClock := structure.NewBallClock(30)
	fmt.Print(ballClock.CalculateDaysUntilResetString())
}
