package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
	}
}
 
