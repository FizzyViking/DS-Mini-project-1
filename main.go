package main

type Fork struct {
	inUse     bool
	timesUsed int
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)
	ch5 := make(chan string)
	ch6 := make(chan string)
	ch7 := make(chan string)
	ch8 := make(chan string) /*
		ch9:= make(chan string)
		ch10:= make(chan string)
		ch11:= make(chan string)
		ch12:= make(chan string)
		ch13 := make(chan string)
		ch14 := make(chan string)
		ch15 := make(chan string)
		ch16 := make(chan string)
		ch17:= make(chan string)
		ch18:= make(chan string)
		ch19:= make(chan string)
		ch20:= make(chan string)*/

	go createPhil(ch1, ch2, ch3, ch4).eat()
	go createPhil(ch5, ch6, ch7, ch8).eat()
	go createFork(ch4, ch3, ch1, ch2).pickup()
	go createFork(ch8, ch7, ch5, ch6).pickup()

	go fork(ch1, ch2)
	go phil(ch2, ch1)

	for {
		// have dinner forever ...
	}
}

func fork(chIN, chOUT chan string) {
	<-chIN
	chOUT <- "go ahead and pick me up"
	//For currently in use here
	<-chIN
}

func phil(chIN, chOUT chan string) {
	chOUT <- "im hungry"
	<-chIN

	chOUT <- "im done eating"
}
