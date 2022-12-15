package d3

import (
	"bufio"
	"os"
)

func CalcSumOfPrioritiesByGroup(filename string) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	count := 0
	result := 0
	set1 := map[byte]bool{}
	set2 := map[byte]bool{}

	for scanner.Scan() {
		line := scanner.Bytes()

		if count%3 == 0 {
			for _, c := range line {
				set2[c] = true
			}
		} else if count%3 == 1 {
			for _, c := range line {
				if _, ok := set1[c]; ok {
					set2[c] = true
				}
			}
		} else if count%3 == 2 {
			for _, c := range line {
				if _, ok := set1[c]; ok {
					result += priorities[c]
					break
				}
			}
		}

		count += 1
		set1 = set2
		set2 = map[byte]bool{}
	}

	return result
}
