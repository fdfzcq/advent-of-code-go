package utils

import (
	"os"
	"strconv"
)

func ReadStringFromFile(year int, day int) string {
	bytes, _ := os.ReadFile("inputs/" + strconv.Itoa(year) + "-" + strconv.Itoa(day) + ".txt")
	return string(bytes)
}

func ReadStringFromTestFile(year int, day int) string {
	bytes, _ := os.ReadFile("inputs/" + strconv.Itoa(year) + "-" + strconv.Itoa(day) + "-test.txt")
	return string(bytes)
}
