package main

import (
	"fmt"
	"os"
	"bufio"
)


func getSharedChar(a string, b string) byte {
	for i := 0; i < len(a); i ++{
		for j := 0; j < len(b); j++{
			if a[i] == b[j]{
				return a[i]
			}
		}
	}
	return 0x00
}

func getPriority(char byte) int {
	if char >= byte(97) {
		return int(char - 96)
	} else {
		return int(27 + char - 65)
	}
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	total := 0

	for scanner.Scan() {
		currLine := scanner.Text()

		if currLine != "" {
			char := getSharedChar( currLine[0: len(currLine)/2], currLine[len(currLine)/2:len(currLine)])
			total += getPriority(char)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "Error: ", err)
		panic(err)
	}
	fmt.Println("Part 1:", total)
}

