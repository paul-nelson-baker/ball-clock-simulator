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
	// Reusable errors that we can specifically check for later, if necessary, by reference
	invalidInputString = errors.New("invalid input: input must be either one number, or a two numbers separated by a single space")
	invalidInputCount  = errors.New("invalid number of inputs: cannot be less than one item or more than two")
	invalidBallCount   = errors.New("invalid ball count: cannot be less than 27 or more than 127")
)

func main() {
	fmt.Println("Instructions: Enter a single number for mode 1, or two numbers separated by a space for mode 2")
	fmt.Println("Mode 1: Ball count only")
	fmt.Println("Mode 2: Ball count and iteration count")
	fmt.Println("Use CTRL+C to exit")
	for {
		fmt.Print("Enter your input here: ")
		if userInputValues, err := getValidUserInput(); err != nil {
			// Anything goes wrong, print out the error and re-ask for input
			fmt.Println(err.Error())
		} else {
			switch len(userInputValues) {
			case 1:
				//If mode 1, calculate how many days before we reset
				ballCount := userInputValues[0]
				fmt.Println("Mode 1::Working with ballcount: " + strconv.Itoa(ballCount))
				fmt.Print(structure.CalculateDaysUntilResetString(ballCount))
			case 2:
				// if mode 2, perform the number of iterations necessary
				// and then print out the state of the clock
				ballCount := userInputValues[0]
				iterations := userInputValues[1]
				fmt.Println("Mode 2::Working with ballcount: " + strconv.Itoa(ballCount))
				fmt.Println("Mode 2::Working with iterations: " + strconv.Itoa(iterations))
				ballClock := structure.NewBallClock(ballCount)
				ballClock.TickMinutes(iterations)
				fmt.Println(ballClock.JsonString())
			}
		}
		fmt.Println()
	}
}

func getValidUserInput() ([]int, error) {
	// Get a single line from the user and split them by space
	reader := bufio.NewReader(os.Stdin)
	userInputLine, _ := reader.ReadString('\n')
	userInputItems := strings.Split(userInputLine, " ")

	// Iterate over the items returned from the user and turn
	// them to integers (regardless of how many)
	var results []int
	for _, value := range userInputItems {
		// If a value isn't a valid numerical string, we short-circuit
		// and return an empty slice along with an error we can print
		if number, err := strconv.Atoi(strings.TrimSpace(value)); err != nil {
			return []int{}, invalidInputString
		} else {
			results = append(results, number)
		}
	}
	// Now we validate we have the correct number of items.
	// We only have mode 1 and 2, so we enforce that via length
	if len(results) < 1 && len(results) > 2 {
		return []int{}, invalidInputCount
	}
	// Now that we have validated the user-date is structurally
	// correct we can enforce that it is also logically correct.
	// Second ball must be in the [27, 127] range. I believe inclusively.
	if results[0] < 27 || results[0] > 127 {
		return []int{}, invalidBallCount
	}
	// We've gone through the gamut of various things that can go
	// wrong, but we've proven the data is sanitary.
	return results, nil
}
