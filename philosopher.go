package main

/*
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

func Eat(phil *Phil) {
	for {
		phil.LchOut <- "am hungry" //Request left fork - Cant request if channel full
		<-phil.LchIn               //Picked up left fork
		phil.RchOut <- "am hungry" //Request right fork -
		<-phil.RchIn               //Picked up Right fork
		//Eating now
		phil.eating = true
		phil.timesAte = phil.timesAte + 1
		phil.LchOut <- "done eating" //Dropping left fork
		phil.RchOut <- "done eating" //Dropping Right fork
	}
}*/
