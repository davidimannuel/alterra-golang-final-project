package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	userId int
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

// GenerateToken jwt ...
func (jwtConf *ConfigJWT) GenerateToken(userId int) (string, error) {
	claims := JwtCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(jwtConf.SecretJWT))
	if err != nil {
		return "", err
	}
	return token, nil
}
