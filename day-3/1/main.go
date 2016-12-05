package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var triangles []string

	file, e := os.Open("../input.txt")
	check(e)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		triangles = append(triangles, scanner.Text())
	}

	var total int
	for i := range triangles {
		if validTriangle(strings.TrimSpace(triangles[i][2:5]),
			strings.TrimSpace(triangles[i][7:10]),
			strings.TrimSpace(triangles[i][12:15])) {
			total++
		}
	}
	fmt.Println(total)

}

func validTriangle(aStr, bStr, cStr string) bool {
	a, e := strconv.Atoi(aStr)
	check(e)
	b, e := strconv.Atoi(bStr)
	check(e)
	c, e := strconv.Atoi(cStr)
	check(e)

	if a+b > c && b+c > a && a+c > b {
		return true
	}
	return false
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}

}
