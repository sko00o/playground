package main

import (
	"bufio"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		if isMale((scan.Text())) {
			os.Stdout.WriteString("IGNORE HIM!")
		} else {
			os.Stdout.WriteString("CHAT WITH HER!")
		}
	}
	if err := scan.Err(); err != nil {
		panic(err)
	}
}

// if the number of distinct characters in one's user name is odd,
// then he is a male, otherwise she is a female.
func isMale(name string) bool {
	bm := int32(0)
	for _, r := range name {
		bm |= 1 << (int32(r-'a') + 1)
	}
	sum := int32(0)
	for bm != 0 {
		sum += bm & 1
		bm >>= 1
	}
	return !(sum%2 == 0)
}
