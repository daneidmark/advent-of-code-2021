package main

import (
	"fmt"

	advent "github.com/daneidmark/advent-of-code-2012"
)

func main() {
	r := advent.FileReader{}
	lines := r.ReadInt("cmd/1/1_a.txt")

	w1 := slidingWindow(1, lines)
	w2 := slidingWindow(3, lines)

	total1 := increasing(w1)
	total2 := increasing(w2)

	fmt.Println(total1)
	fmt.Println(total2)
}

func slidingWindow(window int, v []int) []int {
	slidingWindow := []int{}
	for i := window; i < len(v)+1; i++ {
		slidingWindow = append(slidingWindow, sum(v[i-window:i]))
	}

	return slidingWindow

}

func sum(v []int) int {
	total := 0
	for i := 0; i < len(v); i++ {
		total += v[i]
	}
	return total
}

func increasing(v []int) int {
	total := 0
	for i := 1; i < len(v); i++ {
		if v[i]-v[i-1] > 0 {
			total++
		}
	}
	return total
}
