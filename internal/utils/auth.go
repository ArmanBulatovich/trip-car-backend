package utils

import (
	"crypto/rand"
	"math/big"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/trip/trip-service/internal/auth"
	"github.com/trip/trip-service/internal/models"
)

const (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	specialChars = "!#*"
	allChars     = lowerLetters + upperLetters + digits + specialChars
)

func CreateAdminToken(admin *models.Admin) (string, error) {
	standardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // TODO time.HOUR * 24 -> const
	}
	claims := auth.TokenClaims{
		UserId:         admin.ID,
		UserEmail:      admin.Email,
		UserType:       auth.AdminUserType,
		StandardClaims: standardClaims,
	}
	return auth.GenerateToken(claims)
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func GeneratePassword() (string, error) {
	var password []byte
	for i := 0; i < 10; i++ {
		charIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
		if err != nil {
			return "", err
		}
		password = append(password, allChars[charIndex.Int64()])
	}

	return string(password), nil
}
