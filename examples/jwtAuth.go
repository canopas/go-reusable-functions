package examples

import (
	"fmt"

	"github.com/canopas/go-reusables/jwtAuth"
)

func GenerateTokens() {

	accessKey := "YOUR-JWT-ACCESS-KEY"
	refreshKey := "YOUR-JWT-REFRESH-KEY"

	claims := map[string]interface{}{
		"user_id": 1,
	}

	accessToken := jwtAuth.GenerateAccessToken(accessKey, claims, 24) // 1 day, lifespan in hour ,

	refreshToken := jwtAuth.GenerateRefreshToken(refreshKey, claims, 8760) // 1 year, lifespan in hour

	fmt.Printf("Access token: %s ", &accessToken)
	fmt.Printf("Refresh token: %s ", &refreshToken)
}

func checkValidationsOFTokens(accessToken string, refreshToken string) {
	accessKey := "YOUR-JWT-ACCESS-KEY"
	refreshKey := "YOUR-JWT-REFRESH-KEY"

	claims, err := jwtAuth.ValidateAccessTokenWithData(accessToken, accessKey)

	if err != nil {
		fmt.Errorf("Error when validating accessToken: ", err)
	}

	fmt.Printf("Claims: %v", claims)

	validRefreshToken, err := jwtAuth.ValidateRefreshToken(refreshToken, refreshKey)

	if err != nil {
		fmt.Errorf("Error when validating refreshToken: ", err)
	}

	fmt.Printf("validRefreshToken: %b", validRefreshToken)
}
