package FizzBuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sayFizz(t *testing.T) {
	TwoSpecialNumber(3,5)
	assert.Equal(t, saySlogan(3), "Fizz")
}

func Test_sayBuzz(t *testing.T) {
	TwoSpecialNumber(3,5)
	assert.Equal(t, saySlogan(52), "Buzz")
	assert.Equal(t, saySlogan(10), "Buzz")
}

func Test_sayFizzBuzz(t *testing.T) {
	TwoSpecialNumber(3,5)
	assert.Equal(t, saySlogan(51), "FizzBuzz")
	assert.Equal(t, saySlogan(15), "FizzBuzz")
}

func Test_sayNum(t *testing.T) {
	TwoSpecialNumber(3,5)
	assert.Equal(t, saySlogan(7), "7")
}

func Test_startGame(t *testing.T) {
	TwoSpecialNumber(3,5)
	startGame(100)
}