package d4

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func CalcInclusiveOverlaps(filename string) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d+`)
	result := 0

	for scanner.Scan() {
		nums := re.FindAllString(scanner.Text(), -1)
		l1, _ := strconv.Atoi(nums[0])
		r1, _ := strconv.Atoi(nums[1])
		l2, _ := strconv.Atoi(nums[2])
		r2, _ := strconv.Atoi(nums[3])

		if l1 <= l2 && r1 >= r2 || l1 >= l2 && r1 <= r2 {
			result += 1
		}
	}

	return result
}
