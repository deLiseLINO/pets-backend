package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
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
		logrus.Errorf("failed to generate token %v", err)
		return "", err
	}

	return tokenStr, nil
}
