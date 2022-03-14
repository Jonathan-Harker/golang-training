package testing_exercises

import (
	"fmt"
	"strconv"
)

func printNumberSprint(number int) string {
	return fmt.Sprintf("%d", number)
}

func printNumberFormatInt(number int) string {
	return strconv.FormatInt(int64(number), 10)
}

func printNumberItoa(number int) string {
	return strconv.Itoa(number)
}
