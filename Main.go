package main

import "fmt"

func main() {
	clock := initializeClockMap(30)
	for i := 0; i < 325; i++ {
		tickTockClock(&clock)
	}

	printClock(&clock)
	fmt.Println()
}

func initializeClockMap(ballCount int) map[string][]int {
	clock := map[string][]int{}
	clock["Min"] = []int{}
	clock["FiveMin"] = []int{}
	clock["Hour"] = []int{}
	clock["Main"] = []int{}
	for i := 1; i <= ballCount; i++ {
		clock["Main"] = append(clock["Main"], i)
	}
	return clock
}

func tickTockClock(clock *map[string][]int) {
	var nextBall int
	nextBall, (*clock)["Main"] = (*clock)["Main"][0], (*clock)["Main"][1:]

	(*clock)["Min"] = append((*clock)["Min"], nextBall)
	ballCarryOver := checkBallCarryOver(clock, 5, "Min")
	if ballCarryOver >= 1 {
		(*clock)["FiveMin"] = append((*clock)["FiveMin"], ballCarryOver)
		ballCarryOver = checkBallCarryOver(clock, 12, "FiveMin")
	}
	if ballCarryOver >= 1 {
		(*clock)["Hour"] = append((*clock)["Hour"], ballCarryOver)
		ballCarryOver = checkBallCarryOver(clock, 12, "Hour")
	}
	if ballCarryOver >= 1 {
		(*clock)["Main"] = append((*clock)["Main"], ballCarryOver)
	}

}

func checkBallCarryOver(clock *map[string][]int, threshold int, currentBufferName string) int {
	ballCarryOver := -1
	currentBuffer := (*clock)[currentBufferName]
	length := len(currentBuffer)
	if length == threshold {
		ballCarryOver = currentBuffer[length-1]
		for i := length - 2; i >= 0; i-- {
			(*clock)["Main"] = append((*clock)["Main"], currentBuffer[i])
		}
		(*clock)[currentBufferName] = []int{}
	}
	return ballCarryOver
}

func printClock(clock *map[string][]int) {
	fmt.Printf("%+v\n", clock)
}
