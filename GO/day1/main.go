package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
)

type Elf struct {
	position int
	calories int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	var elfNum, currentCalories int = 1, 0;

	var topN = 3;

	var elves []Elf

	for scanner.Scan(){
		currLine := scanner.Text()

		
		if currLine == "" {	
			elves = append(elves, Elf{elfNum,currentCalories})
			elfNum++
			currentCalories = 0

		} else {
			i, err := strconv.Atoi(currLine)
			if err != nil {
				fmt.Println("bad line, cant parse int.")
				panic(err)
			}
			currentCalories += i;

		}
			
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "Error: ", err)
	}

	if len(elves) < topN{
		fmt.Println("not enough elves to get top: ", topN, " elves: ", len(elves))
	}

	sort.Slice(elves, func (i,j int) bool { return elves[i].calories > elves[j].calories})

	total := 0

	for i := 0; i < topN; i++ {
		total += elves[i].calories
		fmt.Println("Elf:", i + 1, "  ", elves[i])
	
	}

	fmt.Println("Total calories: ", total)
}
 
