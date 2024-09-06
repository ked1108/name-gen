// Copyright 2024-09-05 01:11:04 ked1108. All rights reserved.

package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"strings"
)

const CONS = "bcdfghjklmnpqrstvwxyz"
const VOWS = "aeiou"

func randCons() byte {
	return CONS[rand.IntN(len(CONS))]
}

func randVow() byte {
	return VOWS[rand.IntN(len(VOWS))]
}

func isVow(a byte) bool {
	return strings.Contains(VOWS, string(a))
}

func replaceAtIndex(in string, b byte, i int) string {
	out := []byte(in)
	out[i] = b
	return string(out)
}

func genRandWord(syl int) string {
	op := ""
	for i := 0; i <= syl; i++ {
		switch rand.IntN(4) {
		case 0:
			op += string(rune(randCons()))
			op += string(rune(randVow()))
		case 1:
			op += string(rune(randVow()))
		case 2:
			op += string(rune(randCons()))
			op += string(rune(randVow()))
			op += string(rune(randCons()))
		case 3:
			op += string(rune(randVow()))
			op += string(rune(randCons()))
		}
	}

	return op
}

func countVows(word string) int {
	c := 0
	for _, b := range []byte(word) {
		if isVow(b) {
			c++
		}
	}

	return c
}

func fitness(word string) float64 {
	num_vows := countVows(word)
	length := len(word)
	prob := float64(num_vows) / float64(length)

	return float64(prob)
}

func compare(prob1 float64, prob2 float64) bool {
	return math.Abs(0.5-prob1) < math.Abs(0.5-prob2)
}

func mutateChar(char byte) byte {
	x := char
	if isVow(x) {
		for x == char {
			x = randVow()
		}
	} else {
		for x == char {
			x = randCons()
		}
	}

	return x
}

func mutateWord(word string, prob float64) string {
	op := word
	for i, b := range []byte(word) {
		if rand.Float64() < prob {
			x := mutateChar(b)
			op = replaceAtIndex(op, x, i)
		}
	}

	return op
}

func printWords(words []string, moreWords []string) {
	fmt.Println("GENERATED\t\tMUTATED")
	for i := range len(words) {
		fmt.Printf("%s\t\t%s\n", words[i], moreWords[i])
	}
}

func main() {

	words := []string{}

	for range 10 {
		words = append(words, genRandWord(4))
	}

	mutations := []string{}

	for _, word := range words {
		mutations = append(mutations, mutateWord(word, 0.3))
	}

	printWords(words, mutations)

	fitnesses := []float32{}

	for _, word := range mutations {
		fitnesses = append(fitnesses, float32(fitness(word)))
	}

}
