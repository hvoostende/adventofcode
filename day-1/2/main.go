package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Direction int

const (
	north Direction = iota
	east
	south
	west
	undefined
)

var directions = []string{
	"North",
	"East",
	"South",
	"West",
	"Undefined",
}

func (d Direction) String() string { return directions[d] }

var (
	x, y      int
	locations []string
)

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {
	const input = "L2, L3, L3, L4, R1, R2, L3, R3, R3, L1, L3, R2, R3, L3, R4, R3, R3, L1, L4, R4, L2, R5, R1, L5, R1, R3, L5, R2, L2, R2, R1, L1, L3, L3, R4, R5, R4, L1, L189, L2, R2, L5, R5, R45, L3, R4, R77, L1, R1, R194, R2, L5, L3, L2, L1, R5, L3, L3, L5, L5, L5, R2, L1, L2, L3, R2, R5, R4, L2, R3, R5, L2, L2, R3, L3, L2, L1, L3, R5, R4, R3, R2, L1, R2, L5, R4, L5, L4, R4, L2, R5, L3, L2, R4, L1, L2, R2, R3, L2, L5, R1, R1, R3, R4, R1, R2, R4, R5, L3, L5, L3, L3, R5, R4, R1, L3, R1, L3, R3, R3, R3, L1, R3, R4, L5, L3, L1, L5, L4, R4, R1, L4, R3, R3, R5, R4, R3, R3, L1, L2, R1, L4, L4, L3, L4, L3, L5, R2, R4, L2"
	instructions := strings.Split(input, ", ")

	compas := north

	for _, v := range instructions {

		newCompas, e := changeCompas(compas, v[:1])
		check(e)

		compas = newCompas

		i, err := strconv.Atoi(v[1:])
		check(err)

		if walk(compas, i) {
			break
		}
	}
}

func walk(compas Direction, j int) bool {
	for i := 0; i < j; i++ {
		switch compas {
		case north:
			y++
			if visitedTwice(x, y) {
				return true
			}
		case east:
			x++
			if visitedTwice(x, y) {
				return true
			}
		case south:
			y--
			if visitedTwice(x, y) {
				return true
			}
		case west:
			x--
			if visitedTwice(x, y) {
				return true
			}
		}
	}
	return false

}

func visitedTwice(x, y int) bool {
	sx := strconv.Itoa(x)
	sy := strconv.Itoa(y)
	if checkLocation("x"+sx+"y"+sy, locations) {
		fmt.Println("the first location you visit twice is: ", "x"+sx+"y"+sy)
		if x < 0 {
			x = x * -1
		}
		if y < 0 {
			y = y * -1
		}
		fmt.Printf("This is %v blocks away", x+y)
		return true
	}
	locations = append(locations, "x"+sx+"y"+sy)
	return false
}

func checkLocation(location string, locations []string) bool {
	for i, _ := range locations {
		if location == locations[i] {
			return true
		}
	}
	return false
}

func changeCompas(direction Direction, turn string) (Direction, error) {
	switch direction {
	case north:
		if turn == "L" {
			return west, nil
		} else if turn == "R" {
			return east, nil
		}
	case east:
		if turn == "L" {
			return north, nil
		} else if turn == "R" {
			return south, nil
		}
	case south:
		if turn == "L" {
			return east, nil
		} else if turn == "R" {
			return west, nil
		}
	case west:
		if turn == "L" {
			return south, nil
		} else if turn == "R" {
			return north, nil
		}

	}
	return undefined, errors.New("changeCompas: wrong turn argument")
}

/*
--- Part Two ---
Then, you notice the instructions continue on the back of the Recruiting Document.
Easter Bunny HQ is actually at the first location you visit twice.

For example, if your instructions are R8, R4, R4, R8,
the first location you visit twice is 4 blocks away, due East.

How many blocks away is the first location you visit twice?
*/
