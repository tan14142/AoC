package d7

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CalcTotalSizes(filename string) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	result := 0
	stack := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "$ cd") {
			if strings.HasSuffix(line, "..") {
				cur := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack[len(stack)-1] += cur

				if cur <= 100000 {
					result += cur
				}
			} else {
				stack = append(stack, 0)
			}
		} else if n := int(line[0]); n != 0 {
			prefix := strings.Fields(line)[0]
			size, _ := strconv.Atoi(prefix)
			stack[len(stack)-1] += size
		}
	}

	return result
}
