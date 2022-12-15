package main

import (
	"bufio"
	"os"
	"strconv"
)

func calcMostCalories(filename string) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	result := 0
	total := 0

	for scanner.Scan() {
		cur, err := strconv.Atoi(scanner.Text())

		if err == nil {
			total += cur
		} else {
			if result < total {
				result = total
			}
			total = 0
		}
	}

	return result
}
