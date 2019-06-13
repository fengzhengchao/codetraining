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

func isContainFirstSpecialNum(a int) bool {
	loc := strings.Index(strconv.Itoa(a), strconv.Itoa(firstSpecialNum))
	if loc < 0 {
		return false
	}
	return true
}

func isContainSecondSpecialNum(a int) bool {
	loc := strings.Index(strconv.Itoa(a), strconv.Itoa(secondSpecialNum))
	if loc < 0 {
		return false
	}
	return true
}

func isDividFirstSpecialNum(a int) bool {
	result := a % firstSpecialNum
	if result == 0 {
		return true
	}
	return false
}

func isDividSecondSpecialNum(a int) bool {
	result := a % secondSpecialNum
	if result == 0 {
		return true
	}
	return false
}

func isFizz(a int) bool {
	if isContainFirstSpecialNum(a) == true || isDividFirstSpecialNum(a) == true {
		return true
	}
	return false
}

func isBuzz(a int) bool {
	if isContainSecondSpecialNum(a) == true || isDividSecondSpecialNum(a) == true {
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