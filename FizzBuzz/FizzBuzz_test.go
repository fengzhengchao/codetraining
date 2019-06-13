package FizzBuzz

import (
	"testing"
)

func Test_sayFizz(t *testing.T) {
	TwoSpecialNumber(3,5)
	slogan0 := saySlogan(6)
	slogan1 := saySlogan(13)
	if slogan0 == "Fizz" &&  slogan1 == "Fizz" {
		t.Log("第一个测试通过了")
	}else {
		t.Error("第一个测试没通过")
	}
}

func Test_sayBuzz(t *testing.T) {
	TwoSpecialNumber(3,5)
	slogan0 := saySlogan(52)
	slogan1 := saySlogan(10)
	if  slogan0 == "Buzz" && slogan1 == "Buzz" {
		t.Log("第二个测试通过了")
	}else {
		t.Error("第二个测试没通过")
	}
}

func Test_sayFizzBuzz(t *testing.T) {
	TwoSpecialNumber(3,5)
	slogan0 := saySlogan(51)
	slogan1 := saySlogan(15)
	if slogan0 == "FizzBuzz" && slogan1 == "FizzBuzz" {
		t.Log("第三个测试通过了")
	}else {
		t.Error("第三个测试没通过")
	}
}

func Test_sayNum(t *testing.T) {
	TwoSpecialNumber(3,5)
	slogan := saySlogan(7)
	if  slogan == "7" {
		t.Log("第四个测试通过了")
	}else {
		t.Error("第四个测试没通过")
	}
}

func Test_startGame(t *testing.T) {
	TwoSpecialNumber(3,5)
	startGame(100)
}