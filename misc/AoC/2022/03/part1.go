package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const file = "test.txt"

func main() {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sum := 0
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		l := len(line)
		s, chars := line[:l/2], line[l/2:]
		idx := strings.IndexAny(s, chars)
		c := s[idx]

		if c >= 'A' && c <= 'Z' {
			sum += int(c - 'A' + 27)
		} else if c >= 'a' && c <= 'z' {
			sum += int(c - 'a' + 1)
		}
	}
	if err := scan.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
