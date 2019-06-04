package FizzBuzz

import (
	"testing"
)

func Test_FizzBuzz_1(t *testing.T) {
	TwoSpecialNumber(3,5)
	if slogan := JudgeFirstSpecialNum(1); slogan == true {
		t.Error("测试没通过")
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}

	if slogan := StudentSlogan(14); slogan != "14" {
		t.Error("测试没通过")
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}
}

func Test_FizzBuzz_2(t *testing.T) {
	TwoSpecialNumber(3,5)
	if slogan := StudentSlogan(3); slogan != "Fizz" {
		t.Error("测试没通过")
	} else {
		t.Log("第二个测试通过了") //记录一些你期望记录的信息
	}

	if slogan := StudentSlogan(31); slogan != "Fizz" {
		t.Error("测试没通过")
	} else {
		t.Log("第二个测试通过了") //记录一些你期望记录的信息
	}
}

func Test_FizzBuzz_3(t *testing.T) {
	TwoSpecialNumber(3,5)
	if slogan := StudentSlogan(5); slogan != "Buzz" {
		t.Error("测试没通过")
	} else {
		t.Log("第三个测试通过了") //记录一些你期望记录的信息
	}

	if slogan := StudentSlogan(50); slogan != "Buzz" {
		t.Error("测试没通过")
	} else {
		t.Log("第三个测试通过了") //记录一些你期望记录的信息
	}
}

func Test_FizzBuzz_4(t *testing.T) {
	TwoSpecialNumber(3,5)
	if slogan := StudentSlogan(53); slogan != "FizzBuzz" {
		t.Error("测试没通过")
	} else {
		t.Log("第四个测试通过了") //记录一些你期望记录的信息
	}

	if slogan := StudentSlogan(35); slogan != "FizzBuzz" {
		t.Error("测试没通过")
	} else {
		t.Log("第四个测试通过了") //记录一些你期望记录的信息
	}
}

func Test_StartGame(t *testing.T){
	StartGame(3, 5,100)
}