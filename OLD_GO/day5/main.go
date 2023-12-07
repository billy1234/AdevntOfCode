package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func loadInitialState(input ) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)


	for scanner.Scan() {
		currLine := scanner.Text()

		if currLine != "" {

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "Error: ", err)
		panic(err)
	}

}

