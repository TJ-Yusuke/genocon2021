package main

import (
	"bufio"
	"fmt"
	"os"
)

func attachBase(base string) string {
	switch base {
	case "A":
		return "T"
	case "T":
		return "A"
	case "G":
		return "C"
	case "C":
		return "G"
	default:
		return ""
	}
}

func crateStrand(strand string) string {
	var result string
	for _, base := range strand {
		result = attachBase(string([]rune{base})) + result
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var i int
	for scanner.Scan() {
		if i > 0 {
			in := scanner.Text()
			resultStrand := crateStrand(in)
			fmt.Println(resultStrand)
		}
		i++
	}
}
