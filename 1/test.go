package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// read file
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run test.go <filename>")
	}

	// Get the filename from the command-line arguments
	filename := os.Args[1]

	// Read file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	var left []int
	var right []int

	// parse data
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Fatalf("unexpected line format: %s", line)
		}
		
		// convert strings
		leftVal, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("failed to convert left value: %s", err)
		}
		rightVal, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("failed to convert right value: %s", err)
		}

		// append
		left = append(left, leftVal)
		right = append(right, rightVal)
		
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	
	// sort
	sort.Ints(left)
	sort.Ints(right)

	// calculate distances
	totalDistance := 0
	for i := 0; i < len(left); i++ {
		distance := abs(left[i] - right[i])
		totalDistance += distance
	}
	// output

	fmt.Printf("Total distance: %d\n", totalDistance)


}


func abs(x int) int {
	if x<0 {
		return -x
	}
	return x
}
