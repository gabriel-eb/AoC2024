package day3

import (
	"regexp"
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
			numA, numB := getNumbers(mult)
			sum += numA * numB
		}
	}
	return sum
}

func CompileMultsPart2(fileName string) int {
	lines := utils.ReadFile(fileName)
	sum := 0
	enable := true
	for _, line := range lines {
		mults := getEnabledMults(line, &enable)
		for _, mult := range mults {
			numA, numB := getNumbers(mult)
			sum += numA * numB
		}
	}
	return sum
}

func getNumbers(mult string) (int, int) {
	lastIndex := len(mult) - 1
	strToMult := mult[4:lastIndex]
	numbersToMult := strings.Split(strToMult, ",")
	numA := utils.AtoiCatch(numbersToMult[0])
	numB := utils.AtoiCatch(numbersToMult[1])
	return numA, numB
}

func getEnabledMults(line string, enable *bool) []string {
	rx := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don\'t\(\)`)
	matches := rx.FindAllString(line, len(line))
	mults := []string{}
	for _, match := range matches {
		switch match[:3] {
		case "mul":
			if *enable {
				mults = append(mults, match)
			}
		case "don":
			*enable = false
		case "do(":
			*enable = true
		}
	}
	return mults
}
