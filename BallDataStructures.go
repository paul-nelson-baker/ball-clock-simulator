package main

//import "fmt"
//
//type BallClock struct {
//	registers map[string]BallRegister
//}
//
//type BallRegister struct {
//	name         string
//	balls        []int
//	size         int
//	toNextBuffer chan<- int
//	toMainBuffer chan<- int
//}
//
//func NewBallClock() BallClock {
//	minsBuffer := make(chan int, 5)
//	fiveMinsBuffer := make(chan int, 12)
//	hoursBuffer := make(chan int, 12)
//	mainBuffer := make(chan int, 127)
//
//	minsRegister := NewBallRegister(5, fiveMinsBuffer, mainBuffer)
//	fiveMinsRegister := NewBallRegister(12, hoursBuffer, mainBuffer)
//	hoursRegister := NewBallRegister(12, mainBuffer, mainBuffer)
//	NewBallRegister(127, )
//}
//
//func NewBallRegister(size int, nextBuffer chan<- int, mainBuffer chan<- int) BallRegister {
//	register := BallRegister{
//		balls:        []int{},
//		size:         size,
//		toNextBuffer: nextBuffer,
//		toMainBuffer: mainBuffer,
//	}
//	return register
//}
//
//func (ballRegister *BallRegister) String() string {
//	return fmt.Sprint(ballRegister.balls)
//}