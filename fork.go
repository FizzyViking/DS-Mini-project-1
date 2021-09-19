package main

import (
	"fmt"
	"strconv"
)

type Fork struct {
	name      string
	inUse     bool
	timesUsed int
	LchIn     chan string
	LchOut    chan string
	RchIn     chan string
	RchOut    chan string
	qIn       chan string
	qOut      chan string
}

func createFork(LchIn, LchOut, RchIn, RchOut, qIn, qOut chan string, name string) Fork {
	return Fork{name: name, inUse: false, timesUsed: 0, LchIn: LchIn, LchOut: LchOut, RchIn: RchIn, RchOut: RchOut, qIn: qIn, qOut: qOut}
}

func getQueryF(fork *Fork) {
	for {
		<-fork.qIn
		fmt.Println(fork.name+"| Is in use?: "+strconv.FormatBool(fork.inUse)+" - Times used: ", fork.timesUsed)
		fork.qOut <- "hej"
	}
}

func Pickup(fork *Fork) {
	go getQueryF(fork)
	for {
		select {
		case <-fork.LchIn:
			fork.inUse = true
			fork.timesUsed = fork.timesUsed + 1
			fork.LchOut <- "pick me up"
			<-fork.LchIn
			fork.inUse = false
		case <-fork.RchIn:
			fork.inUse = true
			fork.timesUsed = fork.timesUsed + 1
			fork.RchOut <- "pick me up"
			<-fork.RchIn
			fork.inUse = false
		}
	}
}
