package FizzBuzz

import (
	"fmt"
	"strconv"
)

var firstSpecialNum, secondSpecialNum int

func TwoSpecialNumber(a, b int) {
	firstSpecialNum = a
	secondSpecialNum = b
}

func JudgeFirstSpecialNum(a int) bool {
	if i := a % firstSpecialNum; i == 0 {
		return true
	}
	if j := a % 10; j == firstSpecialNum {
		return true
	}
	if k := a / 10; k == firstSpecialNum {
		return true
	}
	return false
}

func JudgeSecondSpecialNum(a int) bool {
	if i := a % secondSpecialNum; i == 0 {
		return true
	}
	if j := a % 10; j == secondSpecialNum {
		return true
	}
	if k := a / 10; k == secondSpecialNum {
		return true
	}
	return false
}

func StudentSlogan(yourNum int) string {
	if JudgeFirstSpecialNum(yourNum) == true && JudgeSecondSpecialNum(yourNum) == true {
		return "FizzBuzz"
	} else if JudgeFirstSpecialNum(yourNum) == true && JudgeSecondSpecialNum(yourNum) == false {
		return "Fizz"
	} else if JudgeFirstSpecialNum(yourNum) == false && JudgeSecondSpecialNum(yourNum) == true {
		return "Buzz"
	} else {
		return strconv.Itoa(yourNum)
	}
}

func StartGame(a, b, sum int){
	TwoSpecialNumber(a, b)
	for i := 0; i < sum; i++ {
		fmt.Println("student ", i, " say ", StudentSlogan(i))
	}
}
