package main

type BallClock struct {
	mainRegister   BallRegister
	carryRegisters []BallRegister
}

type BallRegister struct {
	balls        []int
	ballIndex    int
	toNextBuffer chan<- int
	toMainBuffer chan<- int
}
