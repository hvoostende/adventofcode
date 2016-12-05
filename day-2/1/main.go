package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	var directions []string

	file, e := os.Open("../input.txt")
	check(e)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		directions = append(directions, scanner.Text())

	}

	var bathroomCode []int

	inputNumber := 5

	for i := range directions {
		for _, w := range directions[i] {
			inputNumber = keypad(w, inputNumber)
		}
		bathroomCode = append(bathroomCode, inputNumber)
	}
	fmt.Println(bathroomCode)
}

func keypad(direction rune, lastNumber int) int {
	switch direction {
	case 'U':
		if lastNumber == 1 || lastNumber == 4 {
			return 1
		}
		if lastNumber == 2 || lastNumber == 5 {
			return 2
		}
		if lastNumber == 3 || lastNumber == 6 {
			return 3
		}
		if lastNumber == 7 {
			return 4
		}
		if lastNumber == 8 {
			return 5
		}
		if lastNumber == 9 {
			return 6
		}
	case 'R':
		if lastNumber == 1 {
			return 2
		}
		if lastNumber == 2 || lastNumber == 3 {
			return 3
		}
		if lastNumber == 4 {
			return 5
		}
		if lastNumber == 5 || lastNumber == 6 {
			return 6
		}
		if lastNumber == 7 {
			return 8
		}
		if lastNumber == 8 || lastNumber == 9 {
			return 9
		}
	case 'D':
		if lastNumber == 1 {
			return 4
		}
		if lastNumber == 2 {
			return 5
		}
		if lastNumber == 3 {
			return 6
		}
		if lastNumber == 4 || lastNumber == 7 {
			return 7
		}
		if lastNumber == 5 || lastNumber == 8 {
			return 8
		}
		if lastNumber == 6 || lastNumber == 9 {
			return 9
		}
	case 'L':
		if lastNumber == 1 || lastNumber == 2 {
			return 1
		}
		if lastNumber == 3 {
			return 2
		}
		if lastNumber == 4 || lastNumber == 5 {
			return 4
		}
		if lastNumber == 6 {
			return 5
		}
		if lastNumber == 7 || lastNumber == 8 {
			return 7
		}
		if lastNumber == 9 {
			return 8
		}
	}
	return 0
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
