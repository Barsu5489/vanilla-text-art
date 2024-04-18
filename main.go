package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Values struct {
	from int
	to   int
}

func main() {
	f, err := os.Open("standard.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := 1

	// For ranges from and to

	// ln := make(map[int]int)
	// for i := 20; i < 127; i++ {
	//
	// }

	var i []int

	for scanner.Scan() {

		line := scanner.Text()

		if len(line) == 0 {
			i = append(i, lines)
		}

		lines++
	}
	fmt.Print((i))
}
