package main

import (
	"fmt"

	d1 "AoC/2022-01"
	d2 "AoC/2022-02"
)

func main() {
	fmt.Println(d1.CalcMostCalories("./2022-01/input.txt"))
	fmt.Println(d1.CalcTop3Calories("./2022-01/input.txt"))
	fmt.Println(d2.CalcTotalScore("./2022-02/input.txt"))
	fmt.Println(d2.CalcTotalScoreUpdated("./2022-02/input.txt"))
}
