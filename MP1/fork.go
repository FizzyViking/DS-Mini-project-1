package main

type Fork struct {
	inUse     bool
	timesUsed int
	LchIn     chan string
	LchOut    chan string
	RchIn     chan string
	RchOut    chan string
}

func createFork(LchIn, LchOut, RchIn, RchOut chan string) Fork {
	return Fork{inUse: false, timesUsed: 0, LchIn: LchIn, LchOut: LchOut, RchIn: RchIn, RchOut: RchOut}
}

func pickup(fork *Fork) {
	lString := <-fork.LchIn //Request from left Phil, all forks must be picked up from left
	if lString == "am hungry" {
		if fork.inUse {
			fork.LchOut <- "busy"
		} else {
			fork.inUse = true
			fork.LchOut <- "pick me up"
		}
	} else if lString == "done eating" {
		fork.inUse = false
	}

	rString := <-fork.RchIn //Request from right Phil, all forks must be picked up from left
	if rString == "am hungry" {
		if fork.inUse {
			fork.RchOut <- "busy"
		} else {
			fork.inUse = true
			fork.RchOut <- "pick me up"
		}
	} else if rString == "done eating" {
		fork.inUse = false
	}

}
