package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24 * 30) // 1 month by develop

	claims := Claims{
		EncodeMD5(username), // It doesn't make sense to encrypt data in JWT token. This is contrary to JWT paradigm.
		EncodeMD5(password), // FIXME: MD5 IS NOT SECURE. Password should not be transferred to frontend part.
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "stskp-api",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenClaims.SignedString(jwtSecret)
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	// FIXME: error not checked, try setup golangci-lint
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// FIXME: tokenClaims.Valid should be checked firstly:
		// when tokenClaims is not null and is not *Claims, this method returns nil, nil
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
