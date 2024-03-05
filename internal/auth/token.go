package auth

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	UserId    uint
	UserEmail string
	UserType  string
	jwt.StandardClaims
}

var AdminUserType = "admin"

func GenerateToken(claims TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("AUTH_TOKEN_KEY")) // TODO: AUTH TOKEN KEY -> env
}

func VerifyToken(token string) *TokenClaims {
	parsedToken, _ := jwt.ParseWithClaims(token, &TokenClaims{}, func(tokenX *jwt.Token) (interface{}, error) {
		return []byte("AUTH_TOKEN_KEY"), nil // TODO: AUTH TOKEN KEY -> env
	})

	return parsedToken.Claims.(*TokenClaims)
}
