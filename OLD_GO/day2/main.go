package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Move int

const (
	ROCK Move = iota + 1
	PAPER 
	SCISSORS
)

const (
	LOSE Move = iota + 1
	TIE
	WIN
)

type Turn struct {
	theirs Move
	mine Move
}

var moveLookup = map[string]Move {
	"A":ROCK,
	"B":PAPER,
	"C":SCISSORS,
	"X":ROCK,
	"Y":PAPER,
	"Z":SCISSORS,
}

const win int = 6

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var turns []Turn
	
	for scanner.Scan() {
		currLine := scanner.Text()

		elems := strings.Split(currLine, " ")

		if len(elems) > 1 {
		

			turns = append(turns ,Turn {moveLookup[elems[0]],moveLookup[elems[1]]})
		}		
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "Error: ", err)
		panic(err)
	}

	totalScorePart1 := 0
	totalScorePart2 := 0

	for _, t := range turns {

		//Part 1 calc
		if t.theirs == t.mine {//Tie
			totalScorePart1 += 3
		} else {
			switch t.theirs {
				case ROCK:
					if t.mine == PAPER { totalScorePart1 += win } //Tie already handled, loss is +0 so no if
				case PAPER:
					if t.mine == SCISSORS { totalScorePart1 += win }
				case SCISSORS:
					if t.mine == ROCK { totalScorePart1 += win }
			}
		}
		totalScorePart1 += int(t.mine)

		//Part 2 calc
		var my_choice Move

		switch t.mine {
			case WIN:
				switch t.theirs {
					case ROCK:
						my_choice = PAPER
					case SCISSORS:
						my_choice = ROCK
					case PAPER:
						my_choice = SCISSORS
				}
				totalScorePart2 += 6
			case TIE:
				my_choice = t.theirs
				totalScorePart2 += 3
			case LOSE:
				switch t.theirs {
					case ROCK:
						my_choice = SCISSORS
					case SCISSORS:
						my_choice = PAPER
					case PAPER:
						my_choice = ROCK
				}
		}

		totalScorePart2 += int(my_choice)



	}

	fmt.Println("Total score for part 1 is: ", totalScorePart1)
	fmt.Println("Total score for part 2 is: ", totalScorePart2)
}
