package main

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

type DayTwo struct{}

const MAX_PROBLEMS = 1

func (d DayTwo) GetDay() int {
  return 2
}

func (d DayTwo) GetInput() interface{} {
  c := make(chan string)
  go ReadFileByLine(2, c)

  res := make([][]int, 0)

  for line := range c {
    levels := strings.Split(line, " ")
    l := make([]int, 0)

    for _, level := range levels {
      num, err := strconv.Atoi(level)
      
      if err != nil {
        panic(err)
      }

      l = append(l, num)
    }

    res = append(res, l)
  }

  return res
}

func (d DayTwo) Test() {
  exampleData := [][]int{
    {7,6,4,2,1},
    {1,2,7,8,9},
    {9,7,6,2,1},
    {1,3,2,4,5},
    {8,6,4,4,1},
    {1,3,6,7,9},
  }

  ans := d.SolveReport(exampleData, true)

  log.Println("test: ", strconv.Itoa(ans))
}

func (d DayTwo) PartOne() {
  data := d.GetInput()

  ans := d.SolveReport(data.([][]int), false)

  log.Println("part one: ", strconv.Itoa(ans))

}

func (d DayTwo) PartTwo() {
  data := d.GetInput()

  ans := d.SolveReport(data.([][]int), true)

  log.Println("part two: ", strconv.Itoa(ans))
}

func (d DayTwo) SolveReport(data [][]int, unsafeLevels bool) int {

  safeReports := 0
  maxProblems := 0

  if unsafeLevels {
    maxProblems = MAX_PROBLEMS
  }

  for idx, level := range data {
    // log.Println(idx, level)

    inc := d.IsIncreasing(level)
    dec := d.IsDecreasing(level)

    if inc == 2 && dec > 2 {
      log.Println("unsafe inc, 2 problems")
      log.Println(idx, level, inc)
    }

    if dec == 2 && inc > 2 {
      log.Println("unsafe dec, 2 problems")
      log.Println(idx, level, dec)
    }

    if inc <= maxProblems {
      safeReports++
      continue
    }

    // if inc := d.IsIncreasing(level); inc <= maxProblems {
    //   safeReports++
    //   continue
    // }

    // if dec := d.IsDecreasing(level); dec <= maxProblems {
    //   safeReports++
    //   continue
    // }
  }

  return safeReports
}

func (d DayTwo) IsIncreasing(level []int) int { 
  prev := level[0]
  problems := 0

  for _, elem := range level[1:] {
    if elem <= prev || elem - prev > 3 {
      // problem with current level, ignore on next iteration
      problems++
      log.Println("problem with elem:", strconv.Itoa(elem), " prev:", strconv.Itoa(prev))
    } else {
      prev = elem
    }
  }
  
  // log.Println("problems: ", strconv.Itoa(problems))
  return problems
}

func (d DayTwo) IsDecreasing(level []int) int {
  c := make([]int, len(level))
  _ = copy(level, c)

  slices.Reverse(c)
  return d.IsIncreasing(c)
}

func (d DayTwo) RemoveElement(level []int, idx int) []int {
  return append(level[:idx], level[idx+1:]...)
}