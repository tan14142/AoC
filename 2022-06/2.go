package d6

import (
	"bufio"
	"os"
)

func CalcStartOfMessage(filename string) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	i := CalcStartOfPacket("./2022-06/input.txt") + 14
	set := map[byte]bool{}

	for j := 0; j < 14; j++ {
		if _, ok := set[line[i-j]]; ok {
			i += 14 - j
			j = -1
			set = map[byte]bool{}
		} else {
			set[line[i-j]] = true
		}
	}

	return i + 1
}
