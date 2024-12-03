package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gabriel-eb/AoC2024/utils"
)

func AnalyzeReports(fileName string) int {
	lines := utils.ReadFile(fileName)
	sum := 0
	for _, line := range lines {
		currLine := strings.Split(line, " ")
		numbers := convertLinesToInts(currLine)
		isSafe := false
		isAscending := numbers[0]-numbers[1] > 0
		for i := 0; i < len(numbers)-1; i++ {
			isSafe = checkSafetiness(numbers[i], numbers[i+1], isAscending)
			if !isSafe {
				break
			}
		}
		if isSafe {
			sum++
		}
	}
	fmt.Printf("Result: %d\n", sum)
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

func AnalyzeReportsPart2(fileName string) int {
	lines := utils.ReadFile(fileName)
	sum := 0
	for _, line := range lines {
		currLine := strings.Split(line, " ")
		numbers := convertLinesToInts(currLine)
		isSafe := false
		checkWithoutIndex(0, numbers, &isSafe)
		for i := 0; i < len(numbers)-1; i++ {
			diff := numbers[i+1] - numbers[i]
			if abs(diff) < 1 || abs(diff) > 3 {
				checkWithoutIndex(i, numbers, &isSafe)
				checkWithoutIndex(i+1, numbers, &isSafe)
				break
			}
			if i+2 < len(numbers) {
				diff2 := numbers[i+2] - numbers[i+1]
				if (diff > 0) != (diff2 > 0) {
					checkWithoutIndex(i, numbers, &isSafe)
					checkWithoutIndex(i+1, numbers, &isSafe)
					checkWithoutIndex(i+2, numbers, &isSafe)
					break
				}
			}
		}
		if isSafe {
			sum++
		}
	}
	fmt.Printf("Result: %d\n", sum)
	return sum
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func checkWithoutIndex(i int, numbers []int, isSafe *bool) {
	temp := append(numbers[:0:0], numbers...)
	temp = append(
		temp[:i], temp[i+1:]...,
	)
	isAscending := temp[0]-temp[1] > 0
	for i := 0; i < len(temp)-1; i++ {
		if !checkSafetiness(temp[i], temp[i+1], isAscending) {
			return
		}
	}
	*isSafe = true
}

func checkSafetiness(numA, numB int, isAscending bool) bool {
	compare := numA - numB
	return compare < 4 &&
		compare > -4 &&
		compare != 0 &&
		isAscending == (compare > 0)
}
