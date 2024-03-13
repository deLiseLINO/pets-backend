package services

import (
	"context"
	"crypto/rand"
	"net/smtp"
	"pets-backend/internal/models"
	"time"

	log "github.com/sirupsen/logrus"
)

type OtpStorage interface {
	Add(
		ctx context.Context,
		code string,
		email string,
		nextSendTime time.Time,
		exparationTime time.Time,
	) error
	GetByEmail(
		ctx context.Context,
		email string,
	) (*models.OTP, error)
}

type OtpService struct {
	otpStorage        OtpStorage
	smtpUser          string
	smtpPassword      string
	smtpHost          string
	smtpPort          string
	otpLength         int
	nextSendPeriodSec int
	exparationTimeMin int
}

func NewOtpService(
	otpStorage OtpStorage,
	smtpUser string,
	smtpPassword string,
	smtpHost string,
	smtpPort string,
	otpLength int,
	nextSendPeriodSec int,
	exparationTimeMin int,
) *OtpService {
	return &OtpService{
		otpStorage:        otpStorage,
		smtpUser:          smtpUser,
		smtpPassword:      smtpPassword,
		smtpHost:          smtpHost,
		smtpPort:          smtpPort,
		otpLength:         otpLength,
		nextSendPeriodSec: nextSendPeriodSec,
		exparationTimeMin: exparationTimeMin,
	}
}

const otpChars = "1234567890"

func (s *OtpService) GenerateOtp() (*models.OTP, error) {
	code, err := s.generateOtp()
	if err != nil {
		log.Error("Failed generating otp: ", err)
		return nil, err
	}

	nextSendTime := time.Now().Add(time.Second * time.Duration(s.nextSendPeriodSec))
	exparationTime := time.Now().Add(time.Minute * time.Duration(s.exparationTimeMin))

	return &models.OTP{
		Code:           code,
		Length:         s.otpLength,
		NextSendTime:   nextSendTime,
		ExparationTime: exparationTime,
	}, nil
}

func (s *OtpService) generateOtp() (string, error) {
	buffer := make([]byte, s.otpLength)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < s.otpLength; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

func (s *OtpService) VerifyCode(ctx context.Context, email string, otpCode string) error {
	otp, err := s.otpStorage.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	if otpCode != otp.Code {
		return models.ErrMismattchCode
	}

	return nil
}

func (s *OtpService) SendOtp(sendingEmail string, otpCode string) error {
	to := []string{sendingEmail}
	message := []byte(otpCode)

	auth := smtp.PlainAuth("", s.smtpUser, s.smtpPassword, s.smtpHost)

	err := smtp.SendMail(s.smtpHost+":"+s.smtpPort, auth, s.smtpUser, to, message)
	if err != nil {
		log.Error("Failed sending otp email: ", err)
		return err
	}

	return nil
}

func (s *OtpService) SaveOtp(
	ctx context.Context,
	code string,
	email string,
	nextSendTime time.Time,
	exparationTime time.Time,
) error {
	return s.otpStorage.Add(ctx, code, email, nextSendTime, exparationTime)
}
