package main

import (
	birdwatcher "exercism/bird_watcher"
	"fmt"
)

func main() {
	fmt.Println(birdwatcher.BirdsInWeek([]int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}, 4))
}
