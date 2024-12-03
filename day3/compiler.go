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
	lines := utils.ReadFile(fileName)
	sum := 0
	enable := true
	for _, line := range lines {
		mults := getEnabledMults(line, &enable)
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

func atoiCatch(str string) int {
	num, err := strconv.Atoi(str)
	utils.Check(err)
	return num
}

func getEnabledMults(line string, enable *bool) []string {
	rx := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don\'t\(\)`)
	matches := rx.FindAllString(line, len(line))
	mults := []string{}
	for _, match := range matches {
		// println(match)
		evaluate := match[:3]
		if evaluate == "don" {
			*enable = false
		}
		if evaluate == "do(" {
			*enable = true
		}
		if *enable && evaluate == "mul" {
			mults = append(mults, match)
		}
	}
	return mults
}
