package examples

import (
	"github.com/canopas/go-reusables/jwtAuth"

	log "github.com/sirupsen/logrus"
)

func GenerateTokens() {

	accessKey := "YOUR-JWT-ACCESS-KEY"
	refreshKey := "YOUR-JWT-REFRESH-KEY"

	claims := map[string]interface{}{
		"user_id": 1,
	}

	accessToken, err := jwtAuth.GenerateAccessToken(accessKey, claims, 24) // 1 day, lifespan in hour ,

	if err != nil{
		log.Error(err)
	}

	refreshToken, err := jwtAuth.GenerateRefreshToken(refreshKey, claims, 8760) // 1 year, lifespan in hour
	if err != nil{
		log.Error(err)
	}

	log.Info("Access token: ", &accessToken)
	log.Info("Refresh token: ", &refreshToken)
}

func checkValidationsOFTokens(accessToken string, refreshToken string) {
	accessKey := "YOUR-JWT-ACCESS-KEY"
	refreshKey := "YOUR-JWT-REFRESH-KEY"

	claims, err := jwtAuth.ValidateAccessTokenWithData(accessToken, accessKey)

	if err != nil {
		log.Error(err)
	}

	log.Info("Claims: ", claims)

	validRefreshToken, err := jwtAuth.ValidateRefreshToken(refreshToken, refreshKey)

	if err != nil {
		log.Error("Error when validating refreshToken: ", err)
	}

	log.Info("validRefreshToken: ", validRefreshToken)
}
