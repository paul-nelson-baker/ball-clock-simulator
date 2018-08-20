package main

import (
	"github.com/paul-nelson-baker/ball-clock-simulator/structure"
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	//fmt.Println(strconv.Atoi(strings.TrimSpace("123\n")))
	fmt.Println("Enter a single number for mode 1, or two numbers separated by a space for mode 2")
	fmt.Println("Mode 1: Ball count only")
	fmt.Println("Mode 2: Ball count and iteration count")
	fmt.Println("Use CTRL+C to exit")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your input here: ")
		userInputLine, _ := reader.ReadString('\n')
		userInputItems := strings.Split(userInputLine, " ")
		//fmt.Println(userInputItems)
		if len(userInputItems) == 1 {
			if ballCount, err := strconv.Atoi(strings.TrimSpace(userInputItems[0])); err == nil && ballCountIsValid(ballCount) {
				fmt.Println("Working with ballcount: " + strconv.Itoa(ballCount))
				ballClock := structure.NewBallClock(ballCount)
				fmt.Print(ballClock.CalculateDaysUntilResetString())
			} else {
				fmt.Println("Only numerical values between 27 and 127 are allowed")
			}
		} else if len(userInputItems) == 2 {
			ballCount, ballCountParseErr := strconv.Atoi(strings.TrimSpace(userInputItems[0]))
			iterations, iterationParseErr := strconv.Atoi(strings.TrimSpace(userInputItems[1]))
			if ballCountParseErr == nil && iterationParseErr == nil && ballCountIsValid(ballCount) {
				fmt.Println("Working with ballcount: " + strconv.Itoa(ballCount))
				fmt.Println("Working with iterations: " + strconv.Itoa(iterations))
				ballClock := structure.NewBallClock(ballCount)
				ballClock.TickMinutes(iterations)
				fmt.Println(ballClock.JsonString())
			} else {
				fmt.Println("Only numerical values between 27 and 127 are allowed")
			}
		} else {
			fmt.Println("Wrong number of inputs, please try again.")
		}
		fmt.Println()
	}
}

func ballCountIsValid(count int) bool {
	return count >= 27 && count <= 127
}
