package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	mathRand "math/rand"
)

func main() {
}

type trader int
type traderArray []trader

// Traders i and j will flip a coin and trade a dollar based on the result.
func (traders traderArray) trade(i, j int) error {
	if traders[i] <= 0 || traders[j] <= 0 {
		return fmt.Errorf("traders %d and %d cannot trade: one or more has zero dollars", i, j)
	}
	flip, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		return fmt.Errorf("could not trade: %w", err)
	}

	switch flip.Uint64() {
	case 0:
		// i won
		traders[i]++
		traders[j]--
	case 1:
		// j won
		traders[j]++
		traders[i]--
	default:
		// I fucked up
		return fmt.Errorf("could not trade: I fucked up my math")
	}

	return nil
}

// Will return a nil slice if no pairs can be made.
func (traders traderArray) makePairs() [][2]int {
	// Contains the index of all non-bankrupt traders.
	var nonBankruptTraders []int
	var pairs [][2]int

	for i, netWorth := range traders {
		if netWorth > 0 {
			nonBankruptTraders = append(nonBankruptTraders, i)
		}
	}

	if nonBankruptTraders == nil || len(nonBankruptTraders) < 2 {
		return nil
	}

	mathRand.Shuffle(
		len(nonBankruptTraders),
		func(i, j int) {
			tmp := nonBankruptTraders[i]
			nonBankruptTraders[i] = nonBankruptTraders[j]
			nonBankruptTraders[j] = tmp
		},
	)

	for i := 0; i < len(nonBankruptTraders)-1; i += 2 {
        pairs = append(pairs, [2]int{nonBankruptTraders[i], nonBankruptTraders[i+1]})
	}

	return pairs
}
