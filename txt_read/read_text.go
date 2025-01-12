package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
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

	fmt.Printf("Array 1 %v\n", left)
	fmt.Printf("Array 2 %v\n", right)

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

		fmt.Printf("Summation: %v ", d)
		dist += d
		fmt.Printf("Distance: %v\n", dist)
	}
}
