package d2

import (
	"bufio"
	"os"
)

func CalcTotalScore(filename string) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	result := 0
	scores := map[string]int{
		"A X": 3 + 1,
		"A Y": 6 + 2,
		"A Z": 0 + 3,
		"B X": 0 + 1,
		"B Y": 3 + 2,
		"B Z": 6 + 3,
		"C X": 6 + 1,
		"C Y": 0 + 2,
		"C Z": 3 + 3,
	}

	for scanner.Scan() {
		result += scores[scanner.Text()]
	}

	return result
}
