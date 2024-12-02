package utils

import (
	"bufio"
	"os"
)

func ReadFile(fileName string) []string {
	file, err := os.Open(fileName)
	Check(err)
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	Check(scanner.Err())

	return lines
}
