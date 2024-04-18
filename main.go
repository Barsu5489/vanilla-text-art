package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var line int
	for scanner.Scan() {
        if line >= 838 && line < 846{
            fmt.Println(scanner.Text())
        }
	
		line++
	}
    fmt.Print(line)
}