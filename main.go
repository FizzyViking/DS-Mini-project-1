package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func queries(phil1, phil2, phil3, phil4, phil5 *Phil, fork1, fork2, fork3, fork4, fork5 *Fork) {
	scanner := bufio.NewScanner(os.Stdin)
	/*
		- each fork must include two channels (one for input and one for

			output, both usable from outside) through which it is possible to

			make queries on the state of the fork (number of times used, in use

			or free)*/
	for true {

		fmt.Println("Choose a philosopher or fork by writting p1-5 or f1-5")
		scanner.Scan()
		cmd := scanner.Text()

		if cmd == "p1" {
			phil1.qIn <- "requesting query"
			phil1.qIn <- "requesting query"

		} else if cmd == "p2" {
			phil2.qIn <- "requesting query"
			phil2.qIn <- "requesting query"

		} else if cmd == "p3" {
			phil3.qIn <- "requesting query"
			phil3.qIn <- "requesting query"

		} else if cmd == "p4" {
			phil4.qIn <- "requesting query"
			phil4.qIn <- "requesting query"

		} else if cmd == "p5" {
			phil5.qIn <- "requesting query"
			phil5.qIn <- "requesting query"

		} else if cmd == "f1" {
			fork1.qIn <- "requesting query"
			fork1.qIn <- "requesting query"

		} else if cmd == "f2" {
			fork2.qIn <- "requesting query"
			fork2.qIn <- "requesting query"

		} else if cmd == "f3" {
			fork3.qIn <- "requesting query"
			fork3.qIn <- "requesting query"

		} else if cmd == "f4" {
			fork4.qIn <- "requesting query"
			fork4.qIn <- "requesting query"

		} else if cmd == "f5" {
			fork5.qIn <- "requesting query"
			fork5.qIn <- "requesting query"
		}
	}

}

func randomPause(max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max*1000)))
}

func main() {
	fmt.Println("Starting main")
	ch1 := make(chan string)
	ch2 := make(chan string)
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

	p1in := make(chan string, 2)
	p1out := make(chan string, 2)
	p2in := make(chan string, 2)
	p2out := make(chan string, 2)
	p3in := make(chan string, 2)
	p3out := make(chan string, 2)
	p4in := make(chan string, 2)
	p4out := make(chan string, 2)
	p5in := make(chan string, 2)
	p5out := make(chan string, 2)

	f1in := make(chan string, 2)
	f1out := make(chan string, 2)
	f2in := make(chan string, 2)
	f2out := make(chan string, 2)
	f3in := make(chan string, 2)
	f3out := make(chan string, 2)
	f4in := make(chan string, 2)
	f4out := make(chan string, 2)
	f5in := make(chan string, 2)
	f5out := make(chan string, 2)

	phil1 := createPhil(ch1, ch2, ch20, ch19, p1in, p1out, "Phil 1")
	phil2 := createPhil(ch5, ch6, ch4, ch3, p2in, p2out, "Phil 2")
	phil3 := createPhil(ch9, ch10, ch8, ch7, p3in, p3out, "Phil 3")
	phil4 := createPhil(ch13, ch14, ch12, ch11, p4in, p4out, "Phil 4")
	phil5 := createPhil(ch17, ch18, ch16, ch15, p5in, p5out, "Phil 5")

	fork1 := createFork(ch3, ch4, ch2, ch1, f1in, f1out, "Fork 1")
	fork2 := createFork(ch7, ch8, ch6, ch5, f2in, f2out, "Fork 2")
	fork3 := createFork(ch11, ch12, ch10, ch9, f3in, f3out, "Fork 3")
	fork4 := createFork(ch15, ch16, ch14, ch13, f4in, f4out, "Fork 4")
	fork5 := createFork(ch19, ch20, ch18, ch17, f5in, f5out, "Fork 5")

	phil1pointer := &phil1
	phil2pointer := &phil2
	phil3pointer := &phil3
	phil4pointer := &phil4
	phil5pointer := &phil5

	fork1pointer := &fork1
	fork2pointer := &fork2
	fork3pointer := &fork3
	fork4pointer := &fork4
	fork5pointer := &fork5

	go EatByStartingLeft(phil1pointer)
	go EatByStartingLeft(phil2pointer)
	go EatByStartingLeft(phil3pointer)
	go EatByStartingLeft(phil4pointer)
	go EatByStartingRight(phil5pointer)

	go Pickup(fork1pointer)
	go Pickup(fork2pointer)
	go Pickup(fork3pointer)
	go Pickup(fork4pointer)
	go Pickup(fork5pointer)

	//go queries(phil1pointer, phil2pointer, phil3pointer, phil4pointer, phil5pointer, fork1pointer, fork2pointer, fork3pointer, fork4pointer, fork5pointer)

	for {
		fmt.Println("==================================")
		// have dinner forever ...
		time.Sleep(time.Millisecond * time.Duration(2*1000))
		p1in <- "requesting query"
		<-p1out
		time.Sleep(time.Millisecond * time.Duration(2*1000))
		p2in <- "requesting query"
		<-p2out
		time.Sleep(time.Millisecond * time.Duration(2*1000))
		p3in <- "requesting query"
		<-p3out
		time.Sleep(time.Millisecond * time.Duration(2*1000))
		p4in <- "requesting query"
		<-p4out
		time.Sleep(time.Millisecond * time.Duration(2*1000))
		p5in <- "requesting query"
		<-p5out

		time.Sleep(time.Millisecond * time.Duration(2*1000))
		f1in <- "requesting query"
		<-f1out
		time.Sleep(time.Millisecond * time.Duration(2*1000))
		f2in <- "requesting query"
		<-f2out
		time.Sleep(time.Millisecond * time.Duration(2*1000))
		f3in <- "requesting query"
		<-f3out
		time.Sleep(time.Millisecond * time.Duration(2*1000))
		f4in <- "requesting query"
		<-f4out
		time.Sleep(time.Millisecond * time.Duration(2*1000))
		f5in <- "requesting query"
		<-f5out
		fmt.Println("------------------")
	}
}
