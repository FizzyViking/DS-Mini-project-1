package main

/*
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

func Pickup(fork *Fork) {

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
*/
