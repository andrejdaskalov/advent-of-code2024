package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFromFile(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("file not present.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputList := make([][]int, 0, 1000)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		splitInts := make([]int, 0, 10)
		for _, val := range split {
			intVal, _ := strconv.Atoi(val)
			splitInts = append(splitInts, intVal)
		}
		inputList = append(inputList, splitInts)
	}

	return inputList
}

func CheckOrder(rep []int) bool {
	if rep[0] == rep[1] {
		return false
	}
	isAsc := rep[0] < rep[1]
	for i := 1; i < len(rep)-1; i++ {
		if isAsc && rep[i] >= rep[i+1] {
			return false
		}
		if !isAsc && rep[i] <= rep[i+1] {
			return false
		}
	}
	return true
}

func absDiff(num1 int, num2 int) int {
	if num1 >= num2 {
		return num1 - num2
	} else {
		return num2 - num1
	}
}

func CheckVariance(rep []int) bool {

	for i := 0; i < len(rep)-1; i++ {
		variance := absDiff(rep[i], rep[i+1])
		if variance < 1 || variance > 3 {
			return false
		}
	}
	return true
}

// Part 1
func CountSafe(input [][]int) int {
	safe := 0
	for _, row := range input {
		order := CheckOrder(row)
		variance := CheckVariance(row)

		if order && variance {
			safe++
		}
	}
	return safe
}

// Part 2
func IsSafe(slc []int) bool {

	for i := range slc {
		new_rep := make([]int, 0, len(slc)-1)
		new_rep = append(new_rep, slc[:i]...)
		new_rep = append(new_rep, slc[i+1:]...)
		if CheckOrder(new_rep) && CheckVariance(new_rep) {
			return true
		}
	}
	return false
}
func CountSafeWithTolerance(slc [][]int) int {
	safe := 0
	for _, row := range slc {
		if IsSafe(row) {
			safe += 1
		}
	}
	return safe
}

func main() {
	input := ReadFromFile("day2-input.txt")

	// safeReports := CountSafe(input)
	// fmt.Printf("Number of safe reports: %d\n", safeReports)

	safeReports := CountSafeWithTolerance(input)
	fmt.Printf("Number of safe reports: %d\n", safeReports)
}
