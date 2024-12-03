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
		numbers := strings.Split(lines[i], " ")
		isSafe := false
		first, err := strconv.Atoi(numbers[0])
		utils.Check(err)
		second, err2 := strconv.Atoi(numbers[1])
		utils.Check(err2)
		isAscending := first-second > 0
		for i := 0; i < len(numbers)-1; i++ {
			num, err := strconv.Atoi(numbers[i])
			utils.Check(err)
			num2, err2 := strconv.Atoi(numbers[i+1])
			utils.Check(err2)
			compare := num - num2
			isSafe = compare < 4 && compare > -4 && compare != 0 && isAscending == (compare > 0)
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
