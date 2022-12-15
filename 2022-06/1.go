package d6

import (
	"bufio"
	"os"
)

func CalcStartOfPacket(filename string) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	i := 3
	set := map[byte]bool{}

	for j := 0; j < 4; j++ {
		if _, ok := set[line[i-j]]; ok {
			i += 4 - j
			j = -1
			set = map[byte]bool{}
		} else {
			set[line[i-j]] = true
		}
	}

	return i + 1
}
