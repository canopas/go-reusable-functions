package util

import (
	random "crypto/rand"
	"errors"
	"io"
)

// generate OTP of given length with numbers only
var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateOTP(length int) (string, error) {

	if length <= 0{
		return "", errors.New("Enter valid length for OTP")
	}

	randomNumberBuf := make([]byte, length)

	_, err := io.ReadAtLeast(random.Reader, randomNumberBuf, length)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(randomNumberBuf); i++ {
		randomNumberBuf[i] = table[int(randomNumberBuf[i])%len(table)]
	}

	return string(randomNumberBuf), nil
}

// Find GCD(greatest common divisor) of given numbers
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Find LCM(Least Common Multiple) of the list
func LCM(list []int) int {
	if len(list) > 0 {
		if len(list) == 1 {
			return list[0]
		}
		if len(list) == 2 {
			return list[0] * list[1] / GCD(list[0], list[1])
		}
		return LCM([]int{list[0], LCM(list[1:])})
	}
	return 0
}