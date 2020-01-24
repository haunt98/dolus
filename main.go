package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	namesSize  = 1024
	maxPostFix = 1000
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("go run main.go names.txt")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("failed to open file", err)
		return
	}

	names := make([]string, 0, namesSize)

	// scan each line then append to names
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		names = append(names, fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		fmt.Println("failed to scan file", err)
		return
	}

	// rand names
	rand.Seed(time.Now().UTC().UnixNano())
	randIndex := rand.Intn(len(names))
	postFix := rand.Intn(maxPostFix)

	fmt.Printf("%s%d\n", names[randIndex], postFix)
}
