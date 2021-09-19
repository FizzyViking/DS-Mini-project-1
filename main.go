package main

import (
	"fmt"
)

type Phil struct {
	name     string
	eating   bool
	timesAte int
	LchIn    chan string
	LchOut   chan string
	RchIn    chan string
	RchOut   chan string
}

func createPhil(LchIn, LchOut, RchIn, RchOut chan string, name string) Phil {
	return Phil{name: name, eating: false, timesAte: 0, LchIn: LchIn, LchOut: LchOut, RchIn: RchIn, RchOut: RchOut}
}

func Eat(phil Phil) {
	for {
		fmt.Println(phil.name + " attempt to request picking up left")
		phil.LchOut <- "am hungry" //Request left fork - Cant request if channel full
		fmt.Println(phil.name + " request has been send left")
		<-phil.LchIn //Picked up left fork
		fmt.Println(phil.name + " confirmed pickup on left")
		fmt.Println(phil.name + " attempt to request picking up right")
		phil.RchOut <- "am hungry" //Request right fork -
		fmt.Println(phil.name + " request has been send right")
		<-phil.RchIn //Picked up Right fork
		fmt.Println(phil.name + " confirmed pickup on right")
		//Eating now
		fmt.Println(phil.name + " is eating::::::::::::::::::::::")
		phil.eating = true
		phil.timesAte = phil.timesAte + 1
		fmt.Println(phil.name + " attempt to drop left")
		phil.LchOut <- "done eating" //Dropping left fork
		fmt.Println(phil.name + " confirmed drop left")
		fmt.Println(phil.name + " attempt to drop Right")
		phil.RchOut <- "done eating" //Dropping Right fork
		fmt.Println(phil.name + " confirmed drop Right")
		fmt.Println(phil.name + " is done eating=============================")
	}
}

func Phil5Eat(phil Phil) {
	for {

		fmt.Println(phil.name + " attempt to request picking up right")
		phil.RchOut <- "am hungry" //Request right fork -
		fmt.Println(phil.name + " request has been send right")
		<-phil.RchIn //Picked up Right fork
		fmt.Println(phil.name + " confirmed pickup on right")
		fmt.Println(phil.name + " attempt to request picking up left")
		phil.LchOut <- "am hungry" //Request left fork - Cant request if channel full
		fmt.Println(phil.name + " request has been send left")
		<-phil.LchIn //Picked up left fork
		fmt.Println(phil.name + " confirmed pickup on left")
		//Eating now
		fmt.Println(phil.name + " is eating::::::::::::::::::::::")
		phil.eating = true
		phil.timesAte = phil.timesAte + 1
		fmt.Println(phil.name + " attempt to drop Right")
		phil.RchOut <- "done eating" //Dropping Right fork
		fmt.Println(phil.name + " confirmed drop Right")
		fmt.Println(phil.name + " attempt to drop left")
		phil.LchOut <- "done eating" //Dropping left fork
		fmt.Println(phil.name + " confirmed drop left")
		fmt.Println(phil.name + " is done eating=============================")
	}
}

type Fork struct {
	name      string
	inUse     bool
	timesUsed int
	LchIn     chan string
	LchOut    chan string
	RchIn     chan string
	RchOut    chan string
}

func createFork(LchIn, LchOut, RchIn, RchOut chan string, name string) Fork {
	return Fork{name: name, inUse: false, timesUsed: 0, LchIn: LchIn, LchOut: LchOut, RchIn: RchIn, RchOut: RchOut}
}

func Pickup(fork Fork) {

	for {
		select {
		case <-fork.LchIn:
			//fmt.Println(fork.name + " fork is picked up by left")
			fmt.Println(fork.name + " in use by right UUUUUUUUUU")
			fork.inUse = true
			fork.timesUsed = fork.timesUsed + 1
			fork.LchOut <- "pick me up"
			<-fork.LchIn
			fmt.Println(fork.name + " not in use by left NNNNNNNNNN")
			//fmt.Println(fork.name + " fork is dropped by left")
			fork.inUse = false
		case <-fork.RchIn:
			//fmt.Println(fork.name + " fork is picked up by right")
			fmt.Println(fork.name + " in use by right UUUUUUUUUU")
			fork.inUse = true
			fork.timesUsed = fork.timesUsed + 1
			fork.RchOut <- "pick me up"
			<-fork.RchIn
			//fmt.Println(fork.name + " fork is dropped by right")
			fmt.Println(fork.name + " not in use by right NNNNNNNNNN")
			fork.inUse = false
		}
	}
}
func send1(in, out chan string) {
	for {
		fmt.Println("starting send")
		out <- "hej"
		<-in
	}

}

func get1(in, out chan string) {
	for {
		fmt.Println("starting get")
		x := <-in
		out <- "ok"
		fmt.Println("In is first " + x)
		x1 := <-in
		fmt.Println("In is then " + x1)
	}
}

func main() {
	fmt.Println("Starting main")
	ch1 := make(chan string)
	ch2 := make(chan string)
	//go send1(ch1, ch2)
	//go get1(ch2, ch1)

	ch3 := make(chan string)
	ch4 := make(chan string)
	ch5 := make(chan string)
	ch6 := make(chan string)
	ch7 := make(chan string)
	ch8 := make(chan string)
	ch9 := make(chan string)
	ch10 := make(chan string)
	ch11 := make(chan string)
	ch12 := make(chan string)
	ch13 := make(chan string)
	ch14 := make(chan string)
	ch15 := make(chan string)
	ch16 := make(chan string)
	ch17 := make(chan string)
	ch18 := make(chan string)
	ch19 := make(chan string)
	ch20 := make(chan string)

	go Eat(createPhil(ch1, ch2, ch20, ch19, "Phil 1"))
	go Eat(createPhil(ch5, ch6, ch4, ch3, "Phil 2"))
	go Eat(createPhil(ch9, ch10, ch8, ch7, "Phil 3"))
	go Eat(createPhil(ch13, ch14, ch12, ch11, "Phil 4"))
	go Phil5Eat(createPhil(ch17, ch18, ch16, ch15, "Phil 5"))
	// Byt om på ch 15 og 18
	go Pickup(createFork(ch3, ch4, ch2, ch1, "Fork 1"))
	go Pickup(createFork(ch7, ch8, ch6, ch5, "Fork 2"))
	go Pickup(createFork(ch11, ch12, ch10, ch9, "Fork 3"))
	go Pickup(createFork(ch15, ch16, ch14, ch13, "Fork 4"))
	go Pickup(createFork(ch19, ch20, ch18, ch17, "Fork 5"))

	for {
		// have dinner forever ...
	}
}
