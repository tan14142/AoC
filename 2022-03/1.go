package d3

import (
	"bufio"
	"os"
)

func CalcSumOfPriorities(filename string) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	priorities := getPriorities()
	result := 0

	for scanner.Scan() {
		set := map[byte]bool{}
		line := scanner.Bytes()
		mid := len(line) / 2

		for i := 0; i < mid; i++ {
			set[line[i]] = true
		}
		for i := mid; i < len(line); i++ {
			if _, ok := set[line[i]]; ok {
				result += priorities[line[i]]
				break
			}
		}
	}

	return result
}
