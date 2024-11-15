package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// loadTriangleFromFile loads a triangle from a JSON file
func loadTriangleFromFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var triangle [][]int
	if err := json.Unmarshal(bytes, &triangle); err != nil {
		return nil, err
	}

	return triangle, nil
}

// max returns the maximum of a and b
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Load the triangle
	triangle, err := loadTriangleFromFile("files/hard.json")
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the maximum path sum
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += max(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	// Print the maximum path sum
	fmt.Println("Maximum path sum:", triangle[0][0])
}
