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

	mapInput := make([][]string, 0, 200)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		mapInput = append(mapInput, split)
	}
	return mapInput
}

type GuardAgent struct {
	positionI    int
	positionJ    int
	direction    string // one of: ^, <, >, v
	mapMatrix    [][]string
	visitedCount int  //distinct positions visited
	inMap        bool //whether the agent is inside the map or not
	firstRight   bool // for part 2
	mvmtPath     [][]int
	obsDir       string
	isLoop       bool
}

func NewGuardAgent(positionI int, positionJ int, direction string, mapMatrix [][]string) GuardAgent {
	return GuardAgent{positionI, positionJ, direction, mapMatrix, 0, true, false, make([][]int, 0), "", false}
}

func (guard *GuardAgent) MoveForward() {

	if guard.mapMatrix[guard.positionI][guard.positionJ] != "X" {
		guard.visitedCount++
		guard.mapMatrix[guard.positionI][guard.positionJ] = "X"
	}

	var nextPosI, nextPosJ int
	if guard.direction == "^" {
		nextPosI = guard.positionI - 1
		nextPosJ = guard.positionJ
	} else if guard.direction == "v" {
		nextPosI = guard.positionI + 1
		nextPosJ = guard.positionJ
	} else if guard.direction == "<" {
		nextPosJ = guard.positionJ - 1
		nextPosI = guard.positionI
	} else if guard.direction == ">" {
		nextPosJ = guard.positionJ + 1
		nextPosI = guard.positionI
	} else {
		// log.Fatal("invalid direction")
	}

	if nextPosI < 0 || nextPosI >= len(guard.mapMatrix) || nextPosJ < 0 || nextPosJ >= len(guard.mapMatrix[0]) {
		guard.inMap = false
		return
	}

	if guard.mapMatrix[nextPosI][nextPosJ] == "#" {
		guard.TurnRight()
		return
	}
	if guard.mapMatrix[nextPosI][nextPosJ] == "O" {
		if guard.obsDir == guard.direction {
			guard.isLoop = true
			guard.visitedCount = -1
			return
		} else {
			guard.obsDir = guard.direction
			guard.TurnRight()
		}
	}
	if guard.firstRight {
		guard.mvmtPath = append(guard.mvmtPath, []int{guard.positionI, guard.positionJ})
	}

	guard.positionI = nextPosI
	guard.positionJ = nextPosJ

}

func (guard *GuardAgent) TurnRight() {
	if guard.direction == "^" {
		guard.direction = ">"
	} else if guard.direction == "v" {
		guard.direction = "<"
	} else if guard.direction == "<" {
		guard.direction = "^"
	} else if guard.direction == ">" {
		guard.direction = "v"
	} else {
		// log.Fatal("invalid direction")
	}
	guard.firstRight = true // for part 2
}

func SimulateGuardMovement(input [][]string, obs []int) (int, [][]int) {
	var guardPosI, guardPosJ int
	var guardDirection string
	for i, row := range input {
		for j, e := range row {
			if e == "^" || e == "<" || e == ">" || e == "v" {
				guardDirection = e
				guardPosI = i
				guardPosJ = j
			}
		}
	}
	mapMatrix := make([][]string, len(input))
	copy(mapMatrix, input)
	if obs != nil {
		mapMatrix[obs[0]][obs[1]] = "O"
	}
	guard := NewGuardAgent(guardPosI, guardPosJ, guardDirection, mapMatrix)
	for guard.inMap && !guard.isLoop {
		fmt.Printf("pos: (%d,%d); dir: %s\n", guard.positionI, guard.positionJ, guard.direction)
		guard.MoveForward()
	}

	return guard.visitedCount, guard.mvmtPath
}

func CheckAllPossibleObstructions(input [][]string) int {
	tmpMap := make([][]string, len(input))
	copy(tmpMap, input)
	_, mvmtPath := SimulateGuardMovement(tmpMap, nil) // initial path
	ch := make(chan bool)
	valid := 0
	for _, pos := range mvmtPath {
		visitedCount, _ := SimulateGuardMovement(input, pos)
		if visitedCount == -1 {
			valid++
		}
		// go func() {
		// 	visitedCount, _ := SimulateGuardMovement(input, pos)
		// 	if visitedCount == -1 {
		// 		ch <- true
		// 	}
		// }()
	}
	for b := range ch {
		if !b {
			continue
		}
		valid++
	}
	return valid
}

func main() {
	input := ReadFromFile("day6-input.txt")
	// count, _ := SimulateGuardMovement(input, nil)
	// fmt.Printf("The guard visited %d distinct positions\n", count)
	possibleObs := CheckAllPossibleObstructions(input)
	fmt.Printf("Possible obstructions: %d\n", possibleObs)
}
