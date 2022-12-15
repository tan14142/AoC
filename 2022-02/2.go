package d2

import (
	"bufio"
	"os"
)

func CalcTotalScoreUpdated(filename string) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	result := 0
	scores := map[string]int{
		"A X": 3 + 0,
		"A Y": 1 + 3,
		"A Z": 2 + 6,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 2 + 0,
		"C Y": 3 + 3,
		"C Z": 1 + 6,
	}

	for scanner.Scan() {
		result += scores[scanner.Text()]
	}

	return result
}
