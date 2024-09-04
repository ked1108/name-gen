// Copyright 2024-09-05 01:11:04 ked1108. All rights reserved.

package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

const CONS = "bcdfghjklmnpqrstvwxyz"
const VOWS = "aeiou"
const SYL = 4

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

func genRandWord() string {
	op := ""
	for i := 0; i <= rand.IntN(SYL-1)+2; i++ {
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

func mutateWord(word string) string {
	op := word
	for i, b := range []byte(word) {
		if isVow(b) {
			op = replaceAtIndex(op, randVow(), i)
		}
	}

	return op
}

func printWords(words []string) {
	for _, word := range words {
		fmt.Println(word)
	}
}

func main() {

	words := []string{}

	for range 10 {
		words = append(words, genRandWord())
	}

	fmt.Println("GENERATED WORDS")
	printWords(words)
	mutations := []string{}

	for _, word := range words {
		mutations = append(mutations, mutateWord(word))
	}

	fmt.Println("\n\n\nMUTATED WORDS")
	printWords(mutations)

}
