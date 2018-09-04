package grcir

import (
	"errors"
	"math"
	"strings"
)

// Votes is a 2d slice of strings. The first dimension
// is each voter and the second is their ranked choices
// where increasing index is a lower priority (0 = first choice, 1
// = 2, etc)
type Votes [][]string

// RunElection takes in a Votes object and returns a winner or set of winners
// in the case of a tie
func RunElection(votes Votes) ([]string, error) {
	// sanity check input
	if len(votes) == 0 {
		return nil, errors.New("no voters")
	}
	numVoters := len(votes)
	fiftyPercent := int(math.Ceil(float64(numVoters) / float64(2)))
	// massage all of the entries to lowercase
	allZero := true
	for _, voterChoices := range votes {
		if allZero && len(voterChoices) > 0 {
			allZero = false
		}
		for index := range voterChoices {
			voterChoices[index] = strings.ToLower(voterChoices[index])
		}
	}
	if allZero {
		return nil, errors.New("voters, but no votes")
	}
	for {
		resultMap := make(map[string]int)
		//tally this round
		for _, voterChoices := range votes {
			currentVotes, ok := resultMap[voterChoices[0]]
			// if there was no entry in the map, this is the first
			// vote for the canditate, set to 1. Otherwise increment
			if !ok {
				resultMap[voterChoices[0]] = 1
			} else {
				resultMap[voterChoices[0]] = currentVotes + 1
			}
		}
		// check if there is a winner or a tie-winner
		winners := make([]string, 0)
		for candidate, numVotes := range resultMap {
			if numVotes >= fiftyPercent {
				winners = append(winners, candidate)
			}
		}
		if len(winners) > 0 {
			return winners, nil
		}
		// if there is no winner, eliminate the lowest vote getters
		// first figure out who to eliminate
		toEliminate := make([]string, 0)
		eliminationCount := int((^uint(0)) >> 1)
		for candidate, numVotes := range resultMap {
			// if it's less than, it's the new person to eliminate,
			// otherwise, it's tied for lowest, and should be appended
			if numVotes < eliminationCount {
				eliminationCount = numVotes
				toEliminate = []string{candidate}
			} else if numVotes == eliminationCount {
				toEliminate = append(toEliminate, candidate)
			}
		}
		// now iterate through and remove anyone whose top choice is toEliminate
		for voterIndex, voterChoices := range votes {
			for _, toEliminate := range toEliminate {
				if voterChoices[0] == toEliminate {
					votes[voterIndex] = voterChoices[1:]
				}
			}
		}
	}
}
