package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type PuzzleInput struct {
	conditions [][]int
	updates    [][]int
}

func ReadFromFile(path string) PuzzleInput {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("file not present.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	conditions := make([][]int, 0, 1500)

	// conditions part
	for scanner.Scan() && scanner.Text() != "" {
		line := scanner.Text()
		split := strings.Split(line, "|")
		cond1, _ := strconv.Atoi(split[0])
		cond2, _ := strconv.Atoi(split[1])
		conditions = append(conditions, []int{cond1, cond2})
	}

	// updates
	updates := make([][]int, 0, 500)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		update := make([]int, 0, 100)
		for _, page := range split {
			pageInt, _ := strconv.Atoi(page)
			update = append(update, pageInt)
		}
		updates = append(updates, update)
	}

	return PuzzleInput{conditions, updates}
}

// Part 1
func CreateMapForUpdate(update []int) map[int]int {

	updateIdxMap := make(map[int]int)
	for i, u := range update {
		updateIdxMap[u] = i
	}

	return updateIdxMap
}

func CheckConditionsForUpdate(input PuzzleInput, updateMap map[int]int) bool {
	for _, cond := range input.conditions {
		idx1, idx1Present := updateMap[cond[0]]
		idx2, idx2Present := updateMap[cond[1]]
		if !idx1Present || !idx2Present {
			continue
		}
		if idx1 >= idx2 {
			return false
		}
	}
	return true
}

func GetSumOfMiddleElements(correctUpdates [][]int) int {
	sum := 0
	for _, update := range correctUpdates {
		sum += update[len(update)/2]
	}
	return sum
}

func GetCorrectUpdates(input PuzzleInput, correct bool) [][]int {
	correctUpdates := make([][]int, 0, len(input.updates))
	for _, update := range input.updates {
		updateMap := CreateMapForUpdate(update)
		isCorrect := CheckConditionsForUpdate(input, updateMap)
		if isCorrect == correct {
			correctUpdates = append(correctUpdates, update)
		}
	}

	return correctUpdates
}

// Part 2
func CorrectUpdate(input PuzzleInput, update []int) {
	corrected := false
	updateMap := CreateMapForUpdate(update)
	for !corrected {
		for _, cond := range input.conditions {
			idx1, idx1Present := updateMap[cond[0]]
			idx2, idx2Present := updateMap[cond[1]]
			if !idx1Present || !idx2Present {
				continue
			}
			if idx1 >= idx2 {
				Swap(update, idx1, idx2)
			}
			updateMap = CreateMapForUpdate(update)
		}

		corrected = CheckConditionsForUpdate(input, updateMap)
	}
}

func Swap(update []int, idx1 int, idx2 int) {
	temp := update[idx1]
	update[idx1] = update[idx2]
	update[idx2] = temp
}

func main() {
	input := ReadFromFile("day5-input.txt")
	// correctUpdates := GetCorrectUpdates(input, true)
	// sum := GetSumOfMiddleElements(correctUpdates)

	incorrectUpdates := GetCorrectUpdates(input, false)
	for _, update := range incorrectUpdates {
		CorrectUpdate(input, update)
	}
	sum := GetSumOfMiddleElements(incorrectUpdates)
	fmt.Printf("Sum of middle elements: %d\n", sum)
}
