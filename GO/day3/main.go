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

func findShared(a string, b string, c string) byte {
	for _, a_char := range a {
		for _, b_char := range b {
			for _, c_char := range c {
				if (a_char == b_char && b_char == c_char) {
					return byte(a_char) //Wont work for characters > 1 byte
				}
			}
		}	
	}
	return 0x00
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	total := 0

	var lines []string	

	for scanner.Scan() {
		currLine := scanner.Text()

		if currLine != "" {
			lines = append(lines, currLine)

			char := getSharedChar( currLine[0: len(currLine)/2], currLine[len(currLine)/2:len(currLine)])
			total += getPriority(char)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "Error: ", err)
		panic(err)
	}

	fmt.Println("Part 1:", total)

	badgeTotal := 0
	
	if len(lines) % 3 == 0 {
		for  i := 0; i < len(lines); i+= 3 {
			badgeTotal += getPriority(findShared(
				lines[i],
				lines[i + 1],
				lines[i + 2]))
		}
	} else{
		fmt.Println(os.Stderr,"Input not divisible by 3, input len: ", len(lines))
	}

	fmt.Println("Part 2:",badgeTotal)

}

