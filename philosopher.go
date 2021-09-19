package main

import (
	"fmt"
	"strconv"
)

type Phil struct {
	name     string
	eating   bool
	timesAte int
	LchIn    chan string
	LchOut   chan string
	RchIn    chan string
	RchOut   chan string
	qIn      chan string
	qOut     chan string
}

func createPhil(LchIn, LchOut, RchIn, RchOut, qIn, qOut chan string, name string) Phil {
	return Phil{name: name, eating: false, timesAte: 0, LchIn: LchIn, LchOut: LchOut, RchIn: RchIn, RchOut: RchOut, qIn: qIn, qOut: qOut}
}

func EatByStartingLeft(phil *Phil) {
	go getQueryP(phil)
	for {
		phil.LchOut <- "am hungry" //Request left fork - Cant request if channel full
		<-phil.LchIn               //Picked up left fork
		phil.RchOut <- "am hungry" //Request right fork -
		<-phil.RchIn               //Picked up Right fork
		phil.eating = true
		phil.timesAte = phil.timesAte + 1
		randomPause(10)
		phil.LchOut <- "done eating" //Dropping left fork
		phil.RchOut <- "done eating" //Dropping Right fork
		phil.eating = false
		//fmt.Println("loop is running")
	}
}
func getQueryP(phil *Phil) {
	for {
		<-phil.qIn
		fmt.Println(phil.name+"| Is eating?: "+strconv.FormatBool(phil.eating)+" - Times eaten: ", phil.timesAte)
		phil.qOut <- "hej"
	}
}

func EatByStartingRight(phil *Phil) {
	go getQueryP(phil)
	for {
		phil.RchOut <- "am hungry" //Request right fork -
		<-phil.RchIn               //Picked up Right fork
		phil.LchOut <- "am hungry" //Request left fork - Cant request if channel full
		<-phil.LchIn               //Picked up left fork
		phil.eating = true
		phil.timesAte = phil.timesAte + 1
		randomPause(10)
		phil.RchOut <- "done eating" //Dropping Right fork
		phil.LchOut <- "done eating" //Dropping left fork
		phil.eating = false
	}
}
