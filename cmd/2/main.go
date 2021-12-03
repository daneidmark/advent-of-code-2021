package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	advent "github.com/daneidmark/advent-of-code-2012"
)

func main() {
	r := advent.FileReader{}
	lines := r.Read("cmd/2/2.txt")

	position := 0
	depth := 0
	aim := 0
	for _, l := range lines {
		a := strings.Split(l, " ")
		direction := a[0]

		x, err := strconv.Atoi(a[1])
		if err != nil {
			log.Fatal("Cannot convert!")
		}
		switch direction {
		case "forward":
			position = position + x
			depth = depth + aim*x
		case "down":
			aim = aim + x
		case "up":
			aim = aim - x
		}
	}

	fmt.Println(position * depth)
}
