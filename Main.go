package main

import (
	"github.com/paul-nelson-baker/ball-clock-simulator/structure"
	"fmt"
)

func main() {
	ballClock := structure.NewBallClock(30)
	fmt.Println(ballClock.String())
	calculation := ballClock.CalculateDaysUntilReset()
	fmt.Println(calculation)
}
