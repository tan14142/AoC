package d3

func getPriorities() map[byte]int {
	var priorities = map[byte]int{}

	for i := 0; i < 26; i++ {
		priorities[byte('a'+i)] = i + 1
	}
	for i := 0; i < 26; i++ {
		priorities[byte('A'+i)] = i + 27
	}

	return priorities
}
