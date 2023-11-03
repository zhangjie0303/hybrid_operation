package main

import (
	"fmt"
	"math/rand"
	"time"
)

var random_min int = 0
var random_max int = 100
var minus_plus_line int = 50 // 加法 or 减法
var words_blank string = "     "
var words_blank_answer string = "   "
var words_blank_mix string = "\t"

func getRandNum() int {
	// 不包含上下限
	return rand.Intn(random_max-random_min-1) + random_min + 1
}

// true: plus, false: minus
func numPlus() bool {
	return getRandNum() > minus_plus_line
}

func numToStr(num int) string {
	return fmt.Sprintf("%2d", num)
}

func numToStrResult(num int) string {
	return fmt.Sprintf("%3d", num)
}

func calcPlus(num1, num2 int, line_count *int, question, answer *string) {
	*question += numToStr(num1) + "+" + numToStr(num2) + "="
	*answer += numToStr(num1) + "+" + numToStr(num2) + "="
	*line_count++

	*answer += numToStrResult(num1 + num2)
	if *line_count < 4 {
		//fmt.Print("   ")
		*question += words_blank
		*answer += words_blank_answer
	} else {
		*question += "\n"
		*answer += "\n"
		*line_count = 0
	}
}

func calcMinus(num1, num2 int, line_count *int, question, answer *string) {
	if num1 < num2 {
		num1, num2 = num2, num1
	}

	*question += numToStr(num1) + "-" + numToStr(num2) + "="
	*answer += numToStr(num1) + "-" + numToStr(num2) + "="
	*line_count++

	*answer += numToStrResult(num1 - num2)
	if *line_count < 4 {
		*question += words_blank
		*answer += words_blank_answer
	} else {
		*question += "\n"
		*answer += "\n"
		*line_count = 0
	}
}

func calcMix(num1, num2, num3 int, line_count *int, question, answer *string) bool {
	if numPlus() {
		first_result := num1 + num2
		first_step := numToStr(num1) + "+" + numToStr(num2)
		if numPlus() {
			*question += first_step + "+" + numToStr(num3) + "="
			*answer += first_step + "+" + numToStr(num3) + "="
			*answer += numToStrResult(first_result + num3)
		} else {
			if first_result-num3 < 0 {
				return false
			}
			*question += first_step + "-" + numToStr(num3) + "="
			*answer += first_step + "-" + numToStr(num3) + "="
			*answer += numToStrResult(first_result - num3)
		}

		*line_count++
		//fmt.Println(" line_count: ", *line_count)
		if *line_count < 3 {
			*question += words_blank_mix
			*answer += words_blank_mix
		} else {
			*question += "\n"
			*answer += "\n"
			*line_count = 0
		}
	} else {
		if num1 < num2 {
			num1, num2 = num2, num1
		}
		first_result := num1 - num2
		first_step := numToStr(num1) + "-" + numToStr(num2)
		if numPlus() {
			*question += first_step + "+" + numToStr(num3) + "="
			*answer += first_step + "+" + numToStr(num3) + "="
			*answer += numToStr(first_result + num3)
		} else {
			if first_result-num3 < 0 {
				return false
			}
			*question += first_step + "-" + numToStr(num3) + "="
			*answer += first_step + "-" + numToStr(num3) + "="
			*answer += numToStrResult(first_result - num3)
		}

		*line_count++
		if *line_count < 3 {
			*question += words_blank_mix
			*answer += words_blank_mix
		} else {
			*question += "\n"
			*answer += "\n"
			*line_count = 0
		}

	}
	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())
	line_count := 0
	other_line_count := 0
	question := ""
	answer := ""
	i := 0

	div_num := 40
	for i < 61 {
		num1 := getRandNum()
		num2 := getRandNum()
		if i < div_num {
			if numPlus() {
				calcPlus(num1, num2, &line_count, &question, &answer)
			} else {
				calcMinus(num1, num2, &line_count, &question, &answer)
			}
		} else {
			if i == div_num {
				question += "\n"
				answer += "\n"
			}
			num3 := getRandNum()
			ret := calcMix(num1, num2, num3, &other_line_count, &question, &answer)
			if !ret {
				continue
			}
		}
		i++
	}
	fmt.Println(question)
	fmt.Println("-------------------------------------------")
	fmt.Println(answer)
}
