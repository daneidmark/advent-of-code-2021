package advent

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type FileReader struct{}

func (r *FileReader) Read(f string) []string {
	file, err := os.Open(f)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func (r *FileReader) ReadInt(f string) []int {
	linesString := r.Read(f)
	var lines []int

	for _, line := range linesString {
		l, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Cannot convert!")
		}
		lines = append(lines, l)
	}

	return lines
}
