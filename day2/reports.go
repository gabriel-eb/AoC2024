package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gabriel-eb/AoC2024/utils"
)

func AnalyzeReports(fileName string) int {
	fmt.Println("--- Day 2: Red-Nosed Reports ---")
	lines := utils.ReadFile(fileName)
	sum := 0
	for i := range lines {
		lines := strings.Split(lines[i], " ")
		numbers := convertLinesToInts(lines)
		isSafe := false
		isAscending := numbers[0]-numbers[1] > 0
		for i := 0; i < len(numbers)-1; i++ {
			compare := numbers[i] - numbers[i+1]
			isSafe = compare < 4 &&
				compare > -4 &&
				compare != 0 &&
				isAscending == (compare > 0)
			if !isSafe {
				break
			}
		}
		if isSafe {
			sum++
		}
	}
	return sum
}

func convertLinesToInts(lines []string) []int {
	numbers := []int{}
	for i := range lines {
		num, err := strconv.Atoi(lines[i])
		utils.Check(err)
		numbers = append(numbers, num)
	}
	return numbers
}
