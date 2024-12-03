package main

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

type DayOne struct{}


// PartOne implements Solution.
func (d DayOne) PartOne() {
	input := d.GetInput()

	total := d.GetDiff(input[0], input[1])

	log.Println("part one: ", strconv.Itoa(total))
}

func (d DayOne) GetDay() int { return 1 }

// PartTwo implements Solution.
func (d DayOne) PartTwo() {
	input := d.GetInput()

	leftInput := input[0]
	rightInput := input[1]

	sim := d.GetSimulatity(leftInput, rightInput)

	log.Println("part two: ", strconv.Itoa(sim))

}

// Test implements Solution.
func (d DayOne) Test() {
	leftInput := []int{ 3,4,2,1,3,3 }
	rightInput := []int{ 4,3,5,3,9,3 }

	total := d.GetSimulatity(leftInput, rightInput)

	println("test: " + strconv.Itoa(total))
}

func (d DayOne) GetDiff(leftInput []int, rightInput []int) int {
	left := d.SortList(leftInput)
	right := d.SortList(rightInput)

	total := 0

	for idx := 0; idx < len(left); idx++ {
		leftItem := left[idx]
		rightItem := right[idx]

		diff := 0

		if leftItem > rightItem {
			diff = leftItem - rightItem
		} else {
			diff = rightItem - leftItem
		}

		total += diff
	}

	return total
}

func (d DayOne) GetSimulatity(leftInput []int, rightInput []int) int {

	score := 0

	for _, left := range leftInput {
		
		times := 0

		for _, right := range rightInput {
			if left == right {
				times++
			}
		}	

		ans := left * times

		score += ans

	}

	return score
}

func (d DayOne) SortList(list []int) []int {
	slices.Sort(list)
	return list
}

// GetInput implements Solution.
func (d DayOne) GetInput() [][]int {
	left := make([]int, 0)
	right := make([]int, 0)

	c := make(chan string)

	go ReadFileByLine(1, c)

	for {
		line, more := <- c

		if !more {
			break
		}

		numbers := strings.Split(line, "   ")
		l, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		
		r, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}

		left = append(left, l)
		right = append(right, r)

	}
	return [][]int{ left, right }

}