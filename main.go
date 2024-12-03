package main

import (
	"log"
	"strconv"
	"time"
)

func main() {
	var day Solution = DayThree{}

	log.Println("Running day " + strconv.Itoa(day.GetDay()))
	start := time.Now()
	day.Test()
	timeTest := time.Now()
	day.PartOne()
	timePartOne := time.Now()
	day.PartTwo()
	timePartTwo := time.Now()

	log.Println("---------")
	log.Println("Times:")	
	log.Println("	- Test: ", timeTest.Sub(start))
	log.Println("	- PartOne: ", timePartOne.Sub(timeTest))
	log.Println("	- PartTwo: ", timePartTwo.Sub(timePartOne))
	log.Println("		- Total: ", timePartTwo.Sub(start))
	log.Println("---------")
}