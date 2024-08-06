package main

import (
	electionday "exercism/election_day"
	"fmt"
)

func main() {
	var finalResults = map[string]int{
		"Mary": 10,
		"John": 51,
	}

	electionday.DecrementVotesOfCandidate(finalResults, "Mary")

	fmt.Print(finalResults)
}
