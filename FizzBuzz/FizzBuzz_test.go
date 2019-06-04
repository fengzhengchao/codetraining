package FizzBuzz

import (
	"testing"
)

// 两个特殊数都不包含
func Test_FizzBuzz_1(t *testing.T) {
	TwoSpecialNumber(3,5)
	if bFirst := JudgeFirstSpecialNum(1); bFirst == true {
		t.Error("测试没通过, number:1")
		return
	}
	if bSecond := JudgeSecondSpecialNum(8); bSecond == true {
		t.Error("测试没通过, number:8")
		return
	}
	if slogan := StudentSlogan(14); slogan != "14" {
		t.Error("测试没通过")
		return
	}
	t.Log("第一个测试通过了")
}

// 包含特殊数3
func Test_FizzBuzz_2(t *testing.T) {
	TwoSpecialNumber(3,5)
	if bFirst := JudgeFirstSpecialNum(12); bFirst != true {
		t.Error("测试没通过, number:12")
		return
	}

	if slogan := StudentSlogan(31); slogan != "Fizz" {
		t.Error("测试没通过")
		return
	}
	t.Log("第二个测试通过了")
}

// 包含特殊数5
func Test_FizzBuzz_3(t *testing.T) {
	TwoSpecialNumber(3,5)
	if bFirst := JudgeFirstSpecialNum(51); bFirst != true {
		t.Error("测试没通过, number:51")
		return
	}

	if slogan := StudentSlogan(50); slogan != "Buzz" {
		t.Error("测试没通过")
		return
	}
	t.Log("第三个测试通过了")
}

// 包含特殊数3和5
func Test_FizzBuzz_4(t *testing.T) {
	TwoSpecialNumber(3,5)
	if bFirst := JudgeFirstSpecialNum(35); bFirst != true {
		t.Error("测试没通过, number:51")
		return
	}
	if bSecond := JudgeSecondSpecialNum(53); bSecond != true {
		t.Error("测试没通过, number:15")
		return
	}

	if slogan := StudentSlogan(35); slogan != "FizzBuzz" {
		t.Error("测试没通过")
		return
	}
	t.Log("第四个测试通过了") //记录一些你期望记录的信息
}

func Test_StartGame(t *testing.T){
	StartGame(3, 5,100)
}