package FizzBuzz

import (
	"strconv"
	"strings"
	"fmt"
)

var firstSpecialNum, secondSpecialNum int

func TwoSpecialNumber(a, b int) {
	firstSpecialNum = a
	secondSpecialNum = b
}

func isContainSpecialNum(a, specialNum int) bool {
	loc := strings.Index(strconv.Itoa(a), strconv.Itoa(specialNum))
	if loc < 0 {
		return false
	}
	return true
}

func isDividSpecialNum(a, specialNum int) bool {
	if a % specialNum == 0 {
		return true
	}
	return false
}

func isFizz(a int) bool {
	if isContainSpecialNum(a, firstSpecialNum) == true || isDividSpecialNum(a, firstSpecialNum) == true {
		return true
	}
	return false
}

func isBuzz(a int) bool {
	if isContainSpecialNum(a, secondSpecialNum) == true || isDividSpecialNum(a, secondSpecialNum) == true {
		return true
	}
	return false
}

func isFizzBuzz(a int) bool {
	if isFizz(a) == true && isBuzz(a) == true {
		return true
	}
	return false
}

func saySlogan(a int) string {
	if isFizzBuzz(a) == true {
		return "FizzBuzz"
	}
	if isFizz(a) == true {
		return "Fizz"
	}
	if isBuzz(a) == true {
		return "Buzz"
	}
	return strconv.Itoa(a)
}

func startGame(sum int) {
	for i := 0; i < sum; i++ {
		fmt.Println("student ", i, " say ", saySlogan(i))
	}
}