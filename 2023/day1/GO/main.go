package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)



func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var total int = 0
	for scanner.Scan() {
		currLine := scanner.Text()

		if currLine != "" {
			var a , b rune = ' ', ' ';
			for _, e := range currLine {
				if ( byte(e) >= 48 && byte(e) <= 57){
					if (a == ' '){
						a = e;
					}else{
						b = e;
					}
				}
			}
			if b == ' ' {
				b = a
			}
			var res string = string([]rune{a,b})
			i, err := strconv.Atoi(res)
			if (err != nil){
				panic(err)
			}			
			total += i
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "Error: ", err)
		panic(err)
	}
	fmt.Println(total)

}

