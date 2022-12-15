package d5

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func CalcTopCrates2(filename string) string {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d+`)
	stacks := getInitState()

	for scanner.Scan() {
		nums := re.FindAllString(scanner.Text(), 3)
		crates, _ := strconv.Atoi(nums[0])
		from, _ := strconv.Atoi(nums[1])
		from -= 1
		to, _ := strconv.Atoi(nums[2])
		to -= 1

		last := len(stacks[from]) - crates
		cur := stacks[from][last:]
		stacks[from] = stacks[from][:last]
		stacks[to] = append(stacks[to], cur...)
	}

	result := ""

	for _, stack := range stacks {
		last := len(stack) - 1
		result += stack[last]
	}

	return result
}
