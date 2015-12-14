package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// Santa struct defines a santa, it's position and map of visited houses
type Santa struct {
	x, y    int
	visited map[string]int
}

// Move changes the santa position according to the direction and updates
// the visited houses map
func (s *Santa) Move(direction byte) {

	// Update position according to the direction character
	switch direction {
	case 94: // ^ (up)
		s.y++
		break
	case 62: // > (right)
		s.x++
		break
	case 118: // v (down)
		s.y--
		break
	case 60: // < (left)
		s.x--
		break
	}

	// Generate a position string
	posString := fmt.Sprintf("%dx%d", s.x, s.y)

	// Update the visited map
	s.visited[posString]++
}

func main() {
	var input []byte

	// Create the part1 & part2 santas
	var p1Santa = Santa{0, 0, make(map[string]int)}
	var p2Santa = Santa{0, 0, make(map[string]int)}
	var p2Robot = Santa{0, 0, make(map[string]int)}

	// Read input.txt for box sizes
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// For each direction...
	for i, direction := range input {

		// Move the part1 santa
		p1Santa.Move(direction)

		// Move the part2 santa on even visits
		// and move the part2 robot on odd visits
		if i%2 == 0 {
			p2Santa.Move(direction)
		} else {
			p2Robot.Move(direction)
		}
	}

	// Combine both part2 santas visited locations
	combinedSanta := Santa{0, 0, map[string]int{}}
	for k, v := range p2Santa.visited {
		combinedSanta.visited[k] += v
	}
	for k, v := range p2Robot.visited {
		combinedSanta.visited[k] += v
	}

	fmt.Println("--- Part 1 ---")
	fmt.Println("Houses visited at least once by Santa:", len(p1Santa.visited))

	fmt.Println("--- Part 2 ---")
	fmt.Println("Houses visited at least once by Santa:", len(p2Santa.visited))
	fmt.Println("Houses visited at least once by Robot:", len(p2Robot.visited))
	fmt.Println("Total houses visited:", len(combinedSanta.visited))
}
