package jwtAuth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var claims = map[string]interface{}{
	"user_id": 1.0,
	"email":   "test@gmail.com",
}

var accessKey = "testAccessSecretKey"
var refreshKey = "testRefreshSecretKey"
var refreshToken = ""
var accessToken = ""
var err error

func TestTokenWithClaims(t *testing.T) {
	asserts := assert.New(t)
	temp := tokenWithClaims(claims, 24)
	asserts.NotEmpty(temp)
}

func TestGenerateAccessToken(t *testing.T) {
	asserts := assert.New(t)
	accessToken, err = GenerateAccessToken(accessKey, claims, 24)
	asserts.NoError(err)
	asserts.NotEmpty(accessToken)
}

func TestGenerateRefreshToken(t *testing.T) {
	asserts := assert.New(t)
	refreshToken, err = GenerateRefreshToken(refreshKey, claims, 48)
	asserts.NoError(err)
	asserts.NotEmpty(refreshToken)
}

func TestValidateRefreshTokenError(t *testing.T) {
	asserts := assert.New(t)
	isValid, err := ValidateRefreshToken("myRefreshToken", refreshKey)
	asserts.Error(err)
	asserts.Equal(false, isValid)
}

func TestValidateRefreshTokenSuccess(t *testing.T) {
	asserts := assert.New(t)
	isValid, err := ValidateRefreshToken(refreshToken, refreshKey)
	asserts.NoError(err)
	asserts.Equal(true, isValid)
}

func TestValidateAccessTokenError(t *testing.T) {
	asserts := assert.New(t)
	data, err := ValidateAccessTokenWithData("myAccessToken", accessKey)
	asserts.NoError(err)
	asserts.Empty(data)
}

func TestValidateAccessTokenSuccess(t *testing.T) {
	asserts := assert.New(t)
	data, err := ValidateAccessTokenWithData(accessToken, accessKey)
	asserts.NoError(err)
	asserts.Equal(claims, data)
}
