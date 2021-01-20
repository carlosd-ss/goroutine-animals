package main

import (
	"goroutines-music/player"
	"sync"
)

func main() {
	// - - - - - - - -
	// 1 2 3 4 5 6 7 8

	drum := player.Drum{
		Tempo: 500,
	}

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go drum.Eagle("x-x---x-x", wg)
	go drum.Dolphin("x----x--x---", wg)
	go drum.Lion("---x-x-x", wg)
	wg.Wait()

}
