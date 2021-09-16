package main

type Phil struct {
	eating   bool
	timesAte int
	LchIn    chan string
	LchOut   chan string
	RchIn    chan string
	RchOut   chan string
}

func createPhil(LchIn, LchOut, RchIn, RchOut chan string) Phil {
	return Phil{eating: false, timesAte: 0, LchIn: LchIn, LchOut: LchOut, RchIn: RchIn, RchOut: RchOut}
}

func eat(phil *Phil) {
	phil.LchOut <- "am hungry" //Request left fork
	phil.RchOut <- "am hungry" //Request right fork
	s := <-phil.LchIn          //Get left fork
	for s == "busy" {          //If fork busy wait until not busy
		s := <-phil.LchIn //Get left fork
	}
	s := <-phil.RchIn // Get right fork, has to wait for left fork
	for s == "busy" { //If fork busy wait until not busy
		s := <-phil.RchIn //Get left fork
	}
	//Eating now
	phil.eating = true
	phil.timesAte = phil.timesAte + 1
	phil.LchOut <- "done eating" //Dropping left fork
	phil.RchOut <- "done eating" //Dropping Right fork
	phil.eat(phil)
}
