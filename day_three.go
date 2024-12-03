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

	ans := d.Run(data.(string), false)

	log.Println("part one: ", strconv.Itoa(ans))
}

// PartTwo implements Solution.
func (d DayThree) PartTwo() {
	data := d.GetInput()

	ans := d.Run(data.(string), true)

	log.Println("part two: ", strconv.Itoa(ans))
}

func (d DayThree) Run(ins string, enableDo bool) int {
	ans := 0
	enabled := true

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
											if enabled {
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
		if enableDo {
			if char == 'd' {
				if ins[idx + 1] == 'o' {
					if ins[idx + 2] == '(' {
						if ins[idx + 3] == ')' {
							enabled = true
						}
					} else if ins[idx + 2] == 'n' {
						if ins[idx + 3] == '\'' {
							if ins[idx + 4] == 't' {
								if ins[idx + 5] == '(' {
									if ins[idx + 6] == ')' {
										enabled = false
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
	// exampleData := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	exampleData := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	ans := d.Run(exampleData, true)

	log.Println("test: ", strconv.Itoa(ans))
}
