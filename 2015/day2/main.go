package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var input []byte
	var boxes [][]byte

	// Read input.txt for box sizes
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Split input by newlines
	boxes = bytes.Split(input, []byte("\n"))

	totalArea := 0
	ribbonLength := 0

	// For each box, try to create a new box object from the dimensions provided
	// and calculate total area and ribbon length
	for _, b := range boxes {
		box, err := NewBox(string(b))
		if err != nil {
			continue
		}

		totalArea += box.area() + box.smallestSideArea()
		ribbonLength += box.ribbonLength() + box.bowLength()
	}

	// Print total wrapping paper & ribbon length needed
	fmt.Println("Total wrapping paper area:", totalArea)
	fmt.Println("Total ribbon length:", ribbonLength)

}
