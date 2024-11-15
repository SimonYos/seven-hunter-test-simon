package main

import (
	"fmt"
	"math"
	"strconv"
)

// Function to decode the input pattern and find the sequence with the minimum sum
func findMinSumSequence(input string) string {
	var minSumSequence string
	minSum := math.MaxInt32

	// Function to generate sequences
	var generateSequences func(index int, currentSequence []int)
	generateSequences = func(index int, currentSequence []int) {
		// If we reach the end of input string, calculate the sum and update min sequence
		if index == len(input) {
			sum := 0
			for _, num := range currentSequence {
				sum += num
			}
			// Convert current sequence to string
			sequenceStr := ""
			for _, num := range currentSequence {
				sequenceStr += strconv.Itoa(num)
			}

			// Check if this sequence has the minimum sum
			if sum < minSum {
				minSum = sum
				minSumSequence = sequenceStr
			}
			return
		}

		// Get the last number in the current sequence
		lastNum := currentSequence[len(currentSequence)-1]

		// Generate next number based on the current character in input
		switch input[index] {
		case 'L':
			for nextNum := lastNum - 1; nextNum >= 0; nextNum-- {
				generateSequences(index+1, append(currentSequence, nextNum))
			}
		case 'R':
			for nextNum := lastNum + 1; nextNum <= 9; nextNum++ {
				generateSequences(index+1, append(currentSequence, nextNum))
			}
		case '=':
			generateSequences(index+1, append(currentSequence, lastNum))
		}
	}

	for start := 0; start <= 9; start++ {
		generateSequences(0, []int{start})
	}

	return minSumSequence
}

func main() {
	// Get the input input from the user
	var input string
	fmt.Print("Enter the input string (e.g., 'LLRR='): ")
	fmt.Scanln(&input)

	// Find the sequence with the minimum sum
	result := findMinSumSequence(input)

	// Print the result
	fmt.Println("The sequence with the minimum sum is:", result)
}
