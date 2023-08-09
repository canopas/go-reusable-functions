package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateOTPError(t *testing.T) {
	asserts := assert.New(t)
	temp, err := GenerateOTP(0)
	asserts.Error(err)
	asserts.Empty(temp)
}

func TestGenerateOTPSuccess(t *testing.T) {
	asserts := assert.New(t)
	temp, err := GenerateOTP(4)
	asserts.NoError(err)
	asserts.NotEmpty(temp)
}

func TestLCM(t *testing.T) {
	asserts := assert.New(t)
	asserts.Equal(40, LCM([]int{4, 8, 20}))
}

func TestGCD(t *testing.T) {
	asserts := assert.New(t)
	asserts.Equal(6, GCD(36, 30))
}
