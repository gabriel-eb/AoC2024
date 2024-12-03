package day3

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gabriel-eb/AoC2024/utils"
)

func CompileMults(fileName string) int {
	lines := utils.ReadFile(fileName)
	sum := 0
	for _, line := range lines {
		rx := regexp.MustCompile(`mul\(\d+,\d+\)`)
		mults := rx.FindAllString(line, len(line))
		for _, mult := range mults {
			lastIndex := len(mult) - 1
			strToMult := mult[4:lastIndex]
			numbersToMult := strings.Split(strToMult, ",")
			numA := atoiCatch(numbersToMult[0])
			numB := atoiCatch(numbersToMult[1])
			sum += numA * numB
		}
	}
	return sum
}

func CompileMultsPart2(fileName string) int {
	return 0
}

func atoiCatch(str string) int {
	num, err := strconv.Atoi(str)
	utils.Check(err)
	return num
}
