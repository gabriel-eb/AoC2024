package main

import (
	"github.com/gabriel-eb/AoC2024/day1"
	"github.com/gabriel-eb/AoC2024/day2"
	"github.com/gabriel-eb/AoC2024/day3"
)

func main() {
	dayOne()
}

func dayOne() {
	println("--- Day 1: Historian Hysteria ---")
	println("--- Part 1 ---")
	day1.Historian("./day1/input")
	println("--- Part 2 ---")
	day1.HistorianPart2("./day1/input")

	println("--- Day 2: Red-Nosed Reports ---")
	println("--- Part 1 ---")
	day2.AnalyzeReports("./day2/input")
	println("--- Part 2 ---")
	day2.AnalyzeReportsPart2("./day2/input")

	println("--- Day 3: Mull It Over ---")
	println("--- Part 1 ---")
	day3.CompileMults("./day3/input")
	println("--- Part 2 ---")
	day3.CompileMultsPart2("./day3/input")
}
