package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	var elfNum, winningElf int = 0, 0;

	var maxCalories, currentCalories int = 0, 0;

	for scanner.Scan(){
		currLine := scanner.Text()

		
		if currLine == ""{
			elfNum++
			currentCalories = 0
		} else {
			i, err := strconv.Atoi(currLine)
			if err != nil {
				fmt.Println("bad line, cant parse int.")
				panic(err)
			}
			currentCalories += i;

			if currentCalories >= maxCalories{
				maxCalories = currentCalories
				winningElf = elfNum
			}
		}
			
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
	}

	fmt.Println("Winner: ",winningElf, " Calories: ", maxCalories);
}
 
