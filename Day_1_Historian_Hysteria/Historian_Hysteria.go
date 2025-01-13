/**
Andres Guerrero Maldonado
Advent of Code Day 1
Commented out code is code that I had an initial take on until
it was later revised/optimized.
**/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Puzzle input from .txt file
	var left []int
	var right []int
	file, ferr := os.Open("input.txt")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(strings.Trim(scanner.Text(), " "), " ")

		leftInt, err := strconv.Atoi(line[0])
		if err != nil {
			panic(fmt.Sprintf("failed to convert left string to int\n\nerr:\n%v\n", err))
		}
		rightInt, err := strconv.Atoi(line[len(line)-1])
		if err != nil {
			panic(fmt.Sprintf("failed to convert left string to int\n\nerr:\n%v\n", err))
		}

		left = append(left, leftInt)
		right = append(right, rightInt)
	}
	fmt.Println(left)
	//Reorganize both arrays into a ascending order then compare
	sort.Ints(left)
	sort.Ints(right)

	//Finding distance between each point
	var dist int
	for range left {
		leftNum := left[0]
		rightNum := right[0]

		left = left[1:]
		right = right[1:]

		d := leftNum - rightNum

		if d < 0 {
			d *= -1
		}

		dist += d
	}

	fmt.Printf("Distance: %v\n", dist)
}
