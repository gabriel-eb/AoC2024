package utils

import "strconv"

func AtoiCatch(str string) int {
	num, err := strconv.Atoi(str)
	Check(err)
	return num
}
