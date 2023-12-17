package main

import (
	"fmt"
	"bufio"
	"os"
)

var numberStrings = []string{
	"one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",
}


func isAtN(x string, pos int) (result int, skip int) {
	for strNum := 0; strNum < len(numberStrings); strNum++ {
		for i := pos; i < len(x); i++ {
			if (numberStrings[strNum][i - pos] != x[i]) {
				break;
			}
		
			if (i - pos == len(numberStrings[strNum]) - 1) {
				return strNum + 1, (i - pos - 1);
			}
		}
	}
	return -1, 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var total1, total2 int = 0, 0
	for scanner.Scan() {
		currLine := scanner.Text()

		if currLine != "" {
			var a , b int = -1, -1;
			var c, d int = -1, -1;
			for i := 0; i < len(currLine); i ++{
				e := currLine[i]
				if ( byte(e) >= 48 && byte(e) <= 57){
					if (a == -1){
						a = int(e - '0');
					}else{
						b = int(e - '0');
					}

					if (c == -1){
						c = int(e - '0')
					}else{
						d = int(e - '0')
					}
				} else if (byte(e) >= 97 && byte(e) <= 122){
					res, skip := isAtN(currLine,i)
					if (res != -1){
						if (c == -1){
							c = res
						} else {
							d = res
						}
					i += skip
					}
					//try to match the strings
				}
			}
			if b == -1 {
				b = a
			}
			if d == -1 {
				d = c
			}
			total1 += (a * 10) + b
			total2 += (c * 10) + d
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ",err)
		panic(err)
	}
	fmt.Println("Part 1: ", total1)
	fmt.Println("Part 2: ", total2)

}

