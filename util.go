package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Solution interface {
	PartOne() 
	Test() 
	PartTwo() 
	GetInput() interface{}
}

const EOF = "done;"

func ReadFileByLine(day int, c chan string) {
	file, err := os.Open("./inputs/" + strconv.Itoa(day))

	if err != nil {
		log.Fatal("err reading day input:", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		c <- scanner.Text()
	}

	c <- EOF

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}