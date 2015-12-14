package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Box defines a box object
type Box struct {
	height int
	width  int
	length int
}

// NewBox (string) creates a new box from a "HxWxL" string
func NewBox(dimensions string) (Box, error) {
	b := Box{}

	dim := strings.Split(dimensions, "x")
	if len(dim) != 3 {
		return b, fmt.Errorf("Invalid dimensions provided: '%s'", dimensions)
	}

	b.width, _ = strconv.Atoi(dim[0])
	b.height, _ = strconv.Atoi(dim[1])
	b.length, _ = strconv.Atoi(dim[2])

	return b, nil
}

func (b Box) area() int {
	return (2 * b.height * b.length) + (2 * b.width * b.length) + (2 * b.height * b.width)
}

func (b Box) smallestSideArea() int {
	sides := []int{b.width, b.height, b.length}
	sort.Ints(sides)

	return sides[0] * sides[1]
}

func (b Box) ribbonLength() int {
	sides := []int{b.width, b.height, b.length}
	sort.Ints(sides)

	return (2 * sides[0]) + (2 * sides[1])
}

func (b Box) bowLength() int {

	return b.width * b.height * b.length
}
