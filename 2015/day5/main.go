package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// isVowel checks if a character/rune is a vowel
func isVowel(c rune) bool {
	var vowels = []rune("aeiou")
	for _, v := range vowels {
		if c == v {
			return true
		}
	}
	return false
}

// vowerCount returns the number of vowels in a string
func vowelCount(s string) int {
	count := 0
	for _, c := range s {
		if isVowel(c) {
			count++
		}
	}
	return count
}

// letterPairs returns how many same-letter pairs there is in a string
func letterPairs(s string) int {
	count := 0
	for i := range s[:len(s)-1] {
		if s[i] == s[i+1] {
			count++
		}
	}
	return count
}

// letterPairsWithSpace returns how many same-letter pairs there is in a string
// where there is a space between them
func letterPairsWithSpace(s string) int {
	count := 0
	for i := range s[:len(s)-2] {
		if s[i] == s[i+2] {
			count++
		}
	}
	return count
}

// letterPairsAnywhere returns how many same-letter pairs there is anywhere in a string
func letterPairsAnywhere(s string) int {
	count := 0
	for i := 0; i < len(s)-1; i++ {
		substr := s[i : i+2]
		for j := i + 2; j < len(s)-1; j++ {
			checkSubstr := s[j : j+2]
			if substr == checkSubstr {
				fmt.Println(s, substr, checkSubstr)
				count++
			}
		}
	}
	return count
}

// containsBadWords checks if a string contains specific "bad" words
func containsBadWords(s string) bool {
	var badWords = []string{"ab", "cd", "pq", "xy"}

	for _, word := range badWords {
		if strings.Contains(s, word) {
			return true
		}
	}

	return false
}

// isNiceString checks if a string is nice
func isNiceString(s string) bool {
	return vowelCount(s) >= 3 && letterPairs(s) >= 1 && !containsBadWords(s)
}

// isNiceStringP2 checks if a string is nice, Part 2
func isNiceStringP2(s string) bool {
	return letterPairsWithSpace(s) >= 1 && letterPairsAnywhere(s) >= 1
}

func main() {
	var input []byte

	var niceStrings []string
	var niceStringsP2 []string

	// Read input.txt for box sizes
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Split input on newlines and loop through it
	for _, s := range bytes.Split(input, []byte("\n")) {
		var str = string(s) // Convert to string

		if len(str) == 0 {
			continue
		}

		if isNiceString(str) {
			niceStrings = append(niceStrings, str)
		}

		if isNiceStringP2(str) {
			niceStringsP2 = append(niceStringsP2, str)
		}

	}

	fmt.Println("---Part 1---")
	fmt.Println("Nice strings:", len(niceStrings))

	fmt.Println("---Part 2---")
	fmt.Println("Nice strings:", len(niceStringsP2))
}
