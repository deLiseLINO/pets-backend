package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type SSOService struct {
	tokenLifespanHours int
	secretKey          string
}

func NewSSOService(tokenLifespanHours int, secretKey string) *SSOService {
	return &SSOService{
		tokenLifespanHours: tokenLifespanHours,
		secretKey:          secretKey,
	}
}

func (s *SSOService) GenerateToken(userid string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userid
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(s.tokenLifespanHours)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		log.Errorf("failed to generate token %v", err)
		return "", err
	}

	return tokenStr, nil
}

func (s *SSOService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("failed to hash password: %v", err)
		return "", err
	}
	return string(hash), nil
}
