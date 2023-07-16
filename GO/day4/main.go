package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type sectionRange struct {
	start int
	end int

}

type pair struct {
	a sectionRange
	b sectionRange
}

func (p *pair) contains() bool {
	return (p.a.start >= p.b.start && p.a.end <= p.b.end) || (p.b.start >= p.a.start && p.b.end <= p.a.end)
}

func (p *pair) intersects() bool {
	return !(p.a.start > p.b.end || p.a.end < p.b.start)
}

func rangeFromStr(str string) *sectionRange {
	elems := strings.Split(str, "-")

	start, err := strconv.Atoi(elems[0])
	if err != nil { 
		fmt.Println(os.Stderr, "Bad start value: ", elems[0])
		panic(err)
	}

	end, err := strconv.Atoi(elems[1])
	if err != nil { 
		fmt.Println(os.Stderr, "Bad end value: ", elems[1])
		panic(err)
	}


	return &sectionRange{start, end}
}

func pairFromStr(str string) *pair {
	ranges := strings.Split(str, ",")
	return &pair{*rangeFromStr(ranges[0]) , *rangeFromStr(ranges[1])}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	containsCount := 0
	intersectsCount := 0

	for scanner.Scan() {
		currLine := scanner.Text()

		if currLine != "" {
			pair := pairFromStr(currLine)

			if pair.intersects() { //if there is no intersection, there is no contains
				intersectsCount ++
				if pair.contains() {
					containsCount ++
				}
			}

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "Error: ", err)
		panic(err)
	}

	fmt.Println("contains: ",containsCount)
	fmt.Println("intersections: ", intersectsCount)
}

