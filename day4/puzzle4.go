package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFromFile(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("file not present.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputList := make([][]string, 0, 200)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		inputList = append(inputList, split)
	}
	return inputList
}

type Item struct {
	i int
	j int
}

type Direction struct {
	i int
	j int
}

// Part 1

func CheckSequence(matrix [][]string, item Item, direction Direction, sequence string) bool {
	if sequence == "" {
		return true
	}
	currentChar := sequence[0:1]

	if item.i < 0 || item.i >= len(matrix) || item.j < 0 || item.j >= len(matrix) {
		return false
	}
	if matrix[item.i][item.j] != currentChar {
		return false
	}
	var nextSeq string
	if len(sequence) == 1 {
		nextSeq = ""
	} else {
		nextSeq = sequence[1:]
	}
	return CheckSequence(matrix, Item{item.i + direction.i, item.j + direction.j}, direction, nextSeq)
}

func CheckCandidate(matrix [][]string, item Item) int {
	if matrix[item.i][item.j] != "X" {
		return 0
	}
	count := 0
	for i := item.i - 1; i < item.i+2; i++ {
		if i < 0 || i >= len(matrix) {
			continue
		}
		for j := item.j - 1; j < item.j+2; j++ {
			if j < 0 || j >= len(matrix[i]) {
				continue
			}

			if matrix[i][j] == "M" {
				isXmas := CheckSequence(matrix, Item{i, j}, Direction{i - item.i, j - item.j}, "MAS")
				if isXmas {
					count++
				}
			}
		}
	}

	return count
}

func ProcessMatrix(matrix [][]string) int {
	countXmas := 0
	for i, r := range matrix {
		for j := range r {
			count := CheckCandidate(matrix, Item{i, j})
			countXmas += count
		}
	}
	return countXmas
}

//Part 2

func CheckCandidateCross(matrix [][]string, item Item) bool {
	if matrix[item.i][item.j] != "A" {
		return false
	}

	if item.i-1 < 0 || item.i+1 >= len(matrix) || item.j-1 < 0 || item.j+1 >= len(matrix) {
		return false
	}

	if (matrix[item.i-1][item.j-1] == "M" && matrix[item.i+1][item.j+1] == "S" || matrix[item.i-1][item.j-1] == "S" && matrix[item.i+1][item.j+1] == "M") &&
		(matrix[item.i-1][item.j+1] == "M" && matrix[item.i+1][item.j-1] == "S" || matrix[item.i-1][item.j+1] == "S" && matrix[item.i+1][item.j-1] == "M") {
		return true

	}

	return false
}

func ProcessMatrixCross(matrix [][]string) int {
	countXmas := 0
	for i, r := range matrix {
		for j := range r {
			isXmas := CheckCandidateCross(matrix, Item{i, j})
			if isXmas {
				countXmas++
			}
		}
	}
	return countXmas
}

func main() {
	input := ReadFromFile("day4-input.txt")
	// xmasCount := ProcessMatrix(input)
	xmasCount := ProcessMatrixCross(input)
	fmt.Printf("Xmas count: %d\n", xmasCount)
}
