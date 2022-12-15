package main

import (
	"bufio"
	"os"
	"strconv"
)

func calcTop3Calories(filename string) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	result := [3]int{0, 0, 0}
	total := 0

	for scanner.Scan() {
		cur, err := strconv.Atoi(scanner.Text())

		if err == nil {
			total += cur
		} else {
			if result[0] < total {
				result[2] = result[1]
				result[1] = result[0]
				result[0] = total
			} else if result[1] < total {
				result[2] = result[1]
				result[1] = total
			} else if result[2] < total {
				result[2] = total
			}
			total = 0
		}
	}

	return result[0] + result[1] + result[2]
}
