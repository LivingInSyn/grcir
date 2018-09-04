package grcir

import (
	"strings"
	"testing"
)

func TestElectionNoTie(t *testing.T) {
	votes := [][]string{
		{"Bob", "Bill", "Sue"},
		{"Sue", "Bob", "Bill"},
		{"Bill", "Sue", "Bob"},
		{"Bob", "Bill", "Sue"},
		{"Sue", "Bob", "Bill"},
	}
	electionResults, err := RunElection(votes)
	if err != nil {
		t.Error(err)
	}
	if len(electionResults) != 1 {
		t.Error("Too Many Winners")
	}
	if electionResults[0] != strings.ToLower("Sue") {
		t.Error("Wrong Winner")
	}
}

func TestElectionTieOneVote(t *testing.T) {
	votes := [][]string{
		{"Bob"},
		{"Sue"},
	}
	electionResults, err := RunElection(votes)
	if err != nil {
		t.Error(err)
	}
	// make sure there are enough winners
	if len(electionResults) <= 1 {
		t.Error("Too Many Winners")
	}
	// see if bob is a winner
	sawbob := false
	sawsue := false
	for _, winner := range electionResults {
		if winner == strings.ToLower("Bob") {
			sawbob = true
		} else if winner == strings.ToLower("Sue") {
			sawsue = true
		}
	}
	if !(sawbob && sawsue) {
		t.Error("missing a tie winner")
	}
}

func TestElectionTieDoubleElim(t *testing.T) {
	votes := [][]string{
		{"Bob", "Sue"},
		{"Sue", "Bob"},
	}
	electionResults, err := RunElection(votes)
	if err != nil {
		t.Error(err)
	}
	// make sure there are enough winners
	if len(electionResults) <= 1 {
		t.Error("Too Many Winners")
	}
	// see if bob is a winner
	sawbob := false
	sawsue := false
	for _, winner := range electionResults {
		if winner == strings.ToLower("Bob") {
			sawbob = true
		} else if winner == strings.ToLower("Sue") {
			sawsue = true
		}
	}
	if !(sawbob && sawsue) {
		t.Error("missing a tie winner")
	}
}

func TestEmpty(t *testing.T) {
	votes := make([][]string, 0)
	_, err := RunElection(votes)
	if err.Error() != "no voters" {
		t.Error("didn't catch no voters")
	}
}
