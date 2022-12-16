package d7

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func CalcNearestMinSize(filename string) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	sizes := []int{}
	stack := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "$ cd") {
			if strings.HasSuffix(line, "..") {
				cur := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack[len(stack)-1] += cur
				sizes = append(sizes, cur)
			} else {
				stack = append(stack, 0)
			}
		} else if n := int(line[0]); n != 0 {
			prefix := strings.Fields(line)[0]
			size, _ := strconv.Atoi(prefix)
			stack[len(stack)-1] += size
		}
	}

	totalSpace := 70000000
	minSpace := 30000000
	freeSpace := totalSpace - stack[0] - stack[1]
	result := math.MaxInt

	for _, size := range sizes {
		if size+freeSpace >= minSpace && size < result {
			result = size
		}
	}

	return result
}
