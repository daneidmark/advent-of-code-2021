package advent

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type FileReader struct{}

func (r *FileReader) ReadInt(f string) []int {
	file, err := os.Open(f)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []int

	for scanner.Scan() {
		l, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Cannot convert!")
		}
		lines = append(lines, l)
	}

	return lines
}
