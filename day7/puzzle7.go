package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	result   int
	operands []int
}

func IntPow(num int, pow int) int {
	return int(math.Pow(float64(num), float64(pow)))
}

func ReadFromFile(path string) []Equation {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("file not present.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	equations := make([]Equation, 0, 900)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		ints := make([]int, 0, len(split)-1)
		res := split[0]
		resInt, _ := strconv.Atoi(res[:len(res)-1])
		for _, e := range split[1:] {
			tmp, _ := strconv.Atoi(e)
			ints = append(ints, tmp)
		}
		equation := Equation{resInt, ints}
		equations = append(equations, equation)
	}
	return equations
}

func EquationIsPossible(eq Equation) bool {
	matr := make([][]int, len(eq.operands))
	matr[0] = make([]int, IntPow(2, len(eq.operands)-1))
	matr[0][0] = eq.operands[0]
	for i := 1; i < len(eq.operands); i++ {
		matr[i] = make([]int, 0, IntPow(2, len(eq.operands)-1))
		op := eq.operands[i]
		for _, e := range matr[i-1] {
			if e == 0 {
				break
			}
			if op+e <= eq.result {
				matr[i] = append(matr[i], e+op)
			}
			if op*e <= eq.result {
				matr[i] = append(matr[i], e*op)
			}
		}
	}

	for _, e := range matr[len(matr)-1] {
		if e == eq.result {
			return true
		}
	}
	return false
}

func CountEquations(input []Equation) int {
	sum := 0
	for _, eq := range input {
		isValid := EquationIsPossible(eq)
		if isValid {
			sum += eq.result
		}
	}
	return sum
}

func digits(num int) int {
	count := 1
	for num/10 != 0 {
		count++
		num /= 10
	}
	return count
}

func ConcatInts(a int, b int) int {
	return a*IntPow(10, digits(b)) + b
}

func GenerateBinaryCombinations(n int, arr []bool, i int, result *[][]bool) {
	if n-1 == i {
		newSlc := make([]bool, 0, len(arr))
		newSlc = append(newSlc, arr...)
		*result = append(*result, newSlc)
		return
	}

	arr[i] = false
	GenerateBinaryCombinations(n, arr, i+1, result)
	arr[i] = true
	GenerateBinaryCombinations(n, arr, i+1, result)
}

func GetPossibleConcatCombinations(slc []int) [][]int {
	binaryCombs := make([][]bool, 0, IntPow(2, len(slc)-1))
	binSlc := make([]bool, len(slc)-1)
	GenerateBinaryCombinations(len(slc)-1, binSlc, 0, &binaryCombs)
	listCombs := make([][]int, 0, IntPow(2, len(slc)-1))
	for _, comb := range binaryCombs {
		combination := make([]int, 0, len(slc))
		concat := slc[0]
		for j := 0; j < len(comb); j++ {
			if comb[j] == true {
				concat = ConcatInts(concat, slc[j+1])
			} else {
				combination = append(combination, concat)
				concat = slc[j+1]
			}
		}
		combination = append(combination, concat)
		listCombs = append(listCombs, combination)
	}
	return listCombs
}

func EquationIsPossibleWithConcat(eq Equation) bool {
	if EquationIsPossible(eq) {
		return true
	}
	possibleCombinations := GetPossibleConcatCombinations(eq.operands)
	for _, comb := range possibleCombinations {
		if EquationIsPossible(Equation{eq.result, comb}) {
			return true
		}
	}
	return false
}
func CountEquationsWithConcat(input []Equation) int {
	sum := 0
	for _, eq := range input {
		isValid := EquationIsPossibleWithConcat(eq)
		if isValid {
			sum += eq.result
		}
	}
	return sum
}

func main() {
	input := ReadFromFile("day7-input.txt")
	// count := CountEquations(input)
	count := CountEquationsWithConcat(input)
	fmt.Printf("Possible equations: %d\n", count)
}
