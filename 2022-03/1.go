package d3

import (
	"bufio"
	"os"
)

var priorities = map[byte]int{}

func SetPriorities() {
	for i := 0; i < 26; i++ {
		priorities[byte('a'+i)] = i + 1
	}
	for i := 0; i < 26; i++ {
		priorities[byte('A'+i)] = i + 27
	}
}

func CalcSumOfPriorities(filename string) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
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
