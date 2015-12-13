package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const (
	asciiOpenParen  = 40
	asciiCloseParen = 41
)

var (
	currentFloor int
)

func init() {
	currentFloor = 0
}

func main() {
	var instructions []byte

	instructions, err := ioutil.ReadFile("instructions.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i, char := range instructions {
		switch char {
		case asciiOpenParen:
			currentFloor++
			break
		case asciiCloseParen:
			currentFloor--
			break
		}

		if currentFloor == -1 {
			fmt.Println("Santa has entered the basement at instruction:", i+1)
		}
	}

	fmt.Println("Santa ended up on floor:", currentFloor, "after following the instructions.")
}
