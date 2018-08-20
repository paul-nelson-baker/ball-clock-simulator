package main

import (
	"github.com/paul-nelson-baker/ball-clock-simulator/structure"
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"errors"
)

var (
	invalidInputString = errors.New("invalid input: input must be either one number, or a two numbers separated by a single space")
	invalidInputCount  = errors.New("invalid number of inputs: cannot be less than one item or more than two")
	invalidBallCount   = errors.New("invalid ball count: cannot be less than 27 or more than 127")
)

func main() {
	//fmt.Println(strconv.Atoi(strings.TrimSpace("123\n")))
	fmt.Println("Enter a single number for mode 1, or two numbers separated by a space for mode 2")
	fmt.Println("Mode 1: Ball count only")
	fmt.Println("Mode 2: Ball count and iteration count")
	fmt.Println("Use CTRL+C to exit")
	for {
		fmt.Print("Enter your input here: ")
		if ints, e := getValidUserInput(); e != nil {
			fmt.Println(e.Error())
		} else {
			switch len(ints) {
			case 1:
				ballCount := ints[0]
				fmt.Println("Working with ballcount: " + strconv.Itoa(ballCount))
				ballClock := structure.NewBallClock(ballCount)
				fmt.Print(ballClock.CalculateDaysUntilResetString())
			case 2:
				ballCount := ints[0]
				iterations := ints[1]
				fmt.Println("Working with ballcount: " + strconv.Itoa(ballCount))
				fmt.Println("Working with iterations: " + strconv.Itoa(iterations))
				ballClock := structure.NewBallClock(ballCount)
				ballClock.TickMinutes(iterations)
				fmt.Println(ballClock.JsonString())
			}
		}
		fmt.Println()
	}
}

func getValidUserInput() ([]int, error) {
	reader := bufio.NewReader(os.Stdin)
	userInputLine, _ := reader.ReadString('\n')
	userInputItems := strings.Split(userInputLine, " ")

	var results []int
	for _, value := range userInputItems {
		if number, err := strconv.Atoi(strings.TrimSpace(value)); err != nil {
			return []int{}, invalidInputString
		} else {
			results = append(results, number)
		}
	}
	if len(results) < 1 && len(results) > 2 {
		return []int{}, invalidInputCount
	}
	if results[0] < 27 || results[0] > 127 {
		return []int{}, invalidBallCount
	}
	return results, nil
}

func ballCountIsValid(count int) bool {
	return count >= 27 && count <= 127
}
