package main

import (
	"fmt"

	d1 "AoC/2022-01"
	d2 "AoC/2022-02"
	d3 "AoC/2022-03"
	d4 "AoC/2022-04"
	d5 "AoC/2022-05"
	d6 "AoC/2022-06"
	d7 "AoC/2022-07"
)

func main() {
	fmt.Println(d1.CalcMostCalories("./2022-01/input.txt"))
	fmt.Println(d1.CalcTop3Calories("./2022-01/input.txt"))
	fmt.Println(d2.CalcTotalScore("./2022-02/input.txt"))
	fmt.Println(d2.CalcTotalScoreUpdated("./2022-02/input.txt"))
	fmt.Println(d3.CalcSumOfPriorities("./2022-03/input.txt"))
	fmt.Println(d3.CalcSumOfPrioritiesByGroup("./2022-03/input.txt"))
	fmt.Println(d4.CalcInclusiveOverlaps("./2022-04/input.txt"))
	fmt.Println(d4.CalcOverlaps("./2022-04/input.txt"))
	fmt.Println(d5.CalcTopCrates("./2022-05/input.txt"))
	fmt.Println(d5.CalcTopCrates2("./2022-05/input.txt"))
	fmt.Println(d6.CalcStartOfPacket("./2022-06/input.txt"))
	fmt.Println(d6.CalcStartOfMessage("./2022-06/input.txt"))
	fmt.Println(d7.CalcTotalSizes("./2022-07/input.txt"))
	fmt.Println(d7.CalcNearestMinSize("./2022-07/input.txt"))
}
