package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func ReadFromFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("file not present.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputList := make([]string, 0, 10)

	for scanner.Scan() {
		line := scanner.Text()
		inputList = append(inputList, line)
	}

	return inputList
}

func FindMults(str string) [][]int {
	result := make([][]int, 0, 100)
	regex := regexp.MustCompile(`mul\(\d+,\d+\)`)
	number_reg := regexp.MustCompile(`\d+`)
	matches := regex.FindAllString(str, -1)
	for _, match := range matches {
		numbers_idx := number_reg.FindAllStringIndex(match, 2)
		if numbers_idx == nil || len(numbers_idx) != 2 {
			log.Fatal("bad parse")
			continue
		}
		number1, _ := strconv.Atoi(match[numbers_idx[0][0]:numbers_idx[0][1]])
		number2, _ := strconv.Atoi(match[numbers_idx[1][0]:numbers_idx[1][1]])
		result = append(result, []int{number1, number2})
	}
	return result
}

// Part 1
func ParseOperations(input []string) [][]int {
	result := make([][]int, 0, 1000)
	for _, s := range input {
		res := FindMults(s)
		result = slices.Concat(result, res)
	}
	return result
}

func CalculateMults(mults [][]int) int {
	sum := 0
	for _, m := range mults {
		mul_res := m[0] * m[1]
		sum += mul_res
	}
	return sum
}

// Part 2
func ParseOperationsWithConditions(input []string) [][]int {
	var full_input strings.Builder
	reg := regexp.MustCompile(`(do\(\))(.*?)(don't\(\))`)
	result := make([][]int, 0, 1000)

	full_input.WriteString("do()")
	for _, s := range input {
		full_input.WriteString(s)
	}
	full_input.WriteString("don't()")

	matches := reg.FindAllString(full_input.String(), -1)
	for _, m := range matches {
		middle_matches := FindMults(m)
		result = slices.Concat(result, middle_matches)
	}
	return result
}

func main() {
	input := ReadFromFile("day3-input.txt")
	// parsed := ParseOperations(input)
	parsed := ParseOperationsWithConditions(input)
	result := CalculateMults(parsed)

	fmt.Printf("Sum of mults is %d\n", result)
}
