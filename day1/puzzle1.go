package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

	list1 := make([]int, 0, 1000)
	list2 := make([]int, 0, 1000)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		num0, _ := strconv.Atoi(split[0])
		num1, _ := strconv.Atoi(split[1])
		list1 = append(list1, num0)
		list2 = append(list2, num1)
	}

	return [][]int{list1, list2}
}

func absDiff(num1 int, num2 int) int {
	if num1 >= num2 {
		return num1 - num2
	} else {
		return num2 - num1
	}
}

// Part 1
func totalDistance() {
	result := ReadFromFile("day1-input.txt")
	list1 := result[0]
	list2 := result[1]

	slices.Sort(list1)
	slices.Sort(list2)

	total_diff := 0
	for i := 0; i < len(list1); i++ {
		diff := absDiff(list1[i], list2[i])
		// fmt.Printf("%d, %d, diff: %d\n", list1[i], list2[i], diff)
		total_diff += diff
	}

	fmt.Printf("The total distance is %d", total_diff)
}

// Part 2
func countOccurences() {
	result := ReadFromFile("day1-input.txt")
	list1 := result[0]
	list2 := result[1]

	occ := make(map[int]int)

	for _, e1 := range list1 {
		for _, e2 := range list2 {
			if e1 == e2 {
				occ[e1] += 1
			}
		}
	}

	similarity := 0
	for k, v := range occ {
		fmt.Printf("key: %d, val: %d; sim: %d\n", k, occ[k], k*v)
		similarity += (k * v)
	}

	fmt.Printf("Similarity: %d\n", similarity)
}

func main() {

	// totalDistance()

	countOccurences()
}
