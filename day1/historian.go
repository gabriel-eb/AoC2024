package day1

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/gabriel-eb/AoC2024/utils"
)

func Historian(fileName string) int {
	fmt.Println("--- Day 1: Historian Hysteria ---")
	lines := utils.ReadFile(fileName)
	lists := getLists(lines)
	sort.Ints(lists[0])
	sort.Ints(lists[1])
	sum := 0
	for i := range lists[0] {
		op := lists[0][i] - lists[1][i]
		if op < 0 {
			op *= -1
		}
		sum += op
	}
	fmt.Printf("Result: %d\n", sum)
	return sum
}

func getLists(lines []string) [2][]int {
	listA := []int{}
	listB := []int{}
	for _, line := range lines {
		rx := regexp.MustCompile(`\d+`)
		res := rx.FindAllString(line, len(line))
		if len(res) > 1 {
			intA, errA := strconv.Atoi(res[0])
			utils.Check(errA)
			listA = append(listA, intA)
			intB, errB := strconv.Atoi(res[1])
			utils.Check(errB)
			listB = append(listB, intB)
		}
	}
	return [2][]int{[]int(listA), []int(listB)}
}

func HistorianPart2(fileName string) int {
	fmt.Println("--- Day 1, Part 2: Historian Hysteria ---")
	lines := utils.ReadFile(fileName)
	lists := getLists(lines)
	// It is not necessary to sort
	similarities := getSimilarityLists(lists)
	sum := 0
	for i := range similarities {
		sum += similarities[i]
	}
	fmt.Printf("Result: %d\n", sum)
	return sum
}

func getSimilarityLists(lists [2][]int) []int {
	similarities := []int{}
	for i := range lists[0] {
		current_value := lists[0][i]
		count := 0
		for j := range lists[1] {
			if current_value == lists[1][j] {
				count++
			}
		}
		similarities = append(similarities, (current_value * count))
	}
	return similarities
}
