package main

import (
	"log"
	"strconv"
)

type DayThree struct{}

// GetDay implements Solution.
func (d DayThree) GetDay() int {
	return 3
}

// GetInput implements Solution.
func (d DayThree) GetInput() interface{} {
	return ReadFileInOneLine(3)
}

// PartOne implements Solution.
func (d DayThree) PartOne() {
	data := d.GetInput()

	ans := d.Run(data.(string))

	log.Println("part one: ", strconv.Itoa(ans))
}

// PartTwo implements Solution.
func (d DayThree) PartTwo() {
	panic("unimplemented")
}

func (d DayThree) Run(ins string) int {
	ans := 0

	for idx, char := range ins {
		if char == 'm' {
			if ins[idx + 1] == 'u' {
				if ins[idx + 2] == 'l' {
					if ins[idx + 3] == '(' {
						if firstNumber, endIdx := d.FindNumber(ins, idx + 4); endIdx != -1 {
							if secondNumber, endIdx := d.FindNumber(ins, endIdx); endIdx != -1 {
								if ins[endIdx-1] == ')' {
									if fst, err := strconv.Atoi(firstNumber); err == nil {
										if snd, err := strconv.Atoi(secondNumber); err == nil {
											ans += fst * snd
										}
									}
								}
							}
						} 
					}
				}
			}
		}
	}

	return ans
}

func (d DayThree) FindNumber(ins string, idx int) (string, int) {
	str := ""
	endIdx := idx

	for i := idx; i < len(ins); i++ {
		if ins[i] == ',' || ins[i] == ')' {
			endIdx = i+1
			break
		}

		if _, err := strconv.Atoi(string(ins[i])); err == nil {
			str += string(ins[i])
		} else {
			return str, -1
		}
	}

	return str, endIdx
}

// Test implements Solution.
func (d DayThree) Test() {
	exampleData := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	ans := d.Run(exampleData)

	log.Println("test: ", strconv.Itoa(ans))
}
