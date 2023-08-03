package jwtAuth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

type JWTData struct {
	jwt.StandardClaims
	CustomClaims map[string]interface{} `json:"custom,omitempty"`
}

/** 
Generate accessToken usin given data

accessKey: Secret access key 
customClaims: your custom claims you want to add in token like user_id, device_id etc...
lifeSpan: accessToken life time in hours
**/

func GenerateAccessToken(accessKey string, customClaims map[string]interface{}, lifeSpan int) (string, error) {

	token, err := tokenWithClaims(customClaims, lifeSpan).SignedString([]byte(accessKey))

	return token, err
}

/** 
Generate refreshToken usin given data

refreshKey: Secret refresh key 
customClaims: your custom claims you want to add in token like user_id, device_id etc...
lifeSpan: accessToken life time in hours
**/

func GenerateRefreshToken(refreshKey string, customClaims map[string]interface{}, lifeSpan int) (string, error) {
    
	token, err := tokenWithClaims(customClaims, lifeSpan).SignedString([]byte(refreshKey))

	return token, err
}

func tokenWithClaims(customClaims map[string]interface{}, lifeSpan int) *jwt.Token{
	claims := JWTData{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(lifeSpan) * time.Hour).Unix(),
		},
		CustomClaims: customClaims,
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

/** 
Validate given token using given key

token: Access Token you want to validate
key: Secret access key 
**/
func ValidateAccessTokenWithData(token string, key string) (map[string]interface{}, error) {

	claims := &JWTData{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		if validationErr, ok := err.(*jwt.ValidationError); ok {
			if validationErr.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, nil
			} else if validationErr.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, nil
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return claims.CustomClaims, err
}

/** 
Validate given refreshToken using given refreshKey

token: Refresh Token you want to validate
key: Secret refresh key 
**/

func ValidateRefreshToken(token string, key string) (bool, error) {

	claims := &JWTData{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	fmt.Println(parsedToken)
	if err != nil || !parsedToken.Valid {
		log.Warn(err)
		return false, err
	}

	return true, nil
}
