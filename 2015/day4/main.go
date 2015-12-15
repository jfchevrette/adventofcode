package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	input = []byte("yzbqklnj")
)

func beginsWithZeroes(data string, count int) bool {
	zeroes := strings.Repeat("0", count)

	if strings.HasPrefix(data, zeroes) {
		return true
	}

	return false
}

func main() {
	var p1Number, p2Number int = 0, 0
	var p1Hash, p2Hash string

	number := 1
	for {
		data := append(input, []byte(strconv.Itoa(number))...)
		hash := fmt.Sprintf("%x", md5.Sum(data))

		if p1Number == 0 && beginsWithZeroes(hash, 5) {
			p1Number = number
			p1Hash = hash
		}

		if p2Number == 0 && beginsWithZeroes(hash, 6) {
			p2Number = number
			p2Hash = hash
		}

		if p1Number != 0 && p2Number != 0 {
			fmt.Printf("Part1 Number: %d, Hash: %s\n", p1Number, p1Hash)
			fmt.Printf("Part2 Number: %d, Hash: %s\n", p2Number, p2Hash)
			os.Exit(0)
		}

		number++
	}

}
