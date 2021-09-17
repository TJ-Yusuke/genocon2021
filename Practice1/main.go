package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var base []string
	for {
		if scanner.Scan() {
			in := scanner.Text()
			base = append(base, in)
		} else {
			break
		}
	}
	fmt.Println(base)
}
