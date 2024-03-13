package postgres

import (
	"context"
	"pets-backend/internal/ent"
	"pets-backend/internal/ent/otpcodes"
	"pets-backend/internal/models"
	"time"

	"github.com/google/uuid"
)

type OtpStorage struct {
	client *ent.Client
}

func NewOtpStorage(client *ent.Client) *OtpStorage {
	return &OtpStorage{client: client}
}

func (s *OtpStorage) Add(
	ctx context.Context,
	code string,
	email string,
	nextSendTime time.Time,
	exparationTime time.Time,
) error {
	_, err := s.client.OtpCodes.
		Create().
		SetID(uuid.New()).
		SetCode(code).
		SetEmail(email).
		SetNextSendTime(nextSendTime).
		SetExparationTime(exparationTime).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *OtpStorage) GetByEmail(
	ctx context.Context,
	email string,
) (*models.OTP, error) {
	otp, err := s.client.OtpCodes.
		Query().
		Where(
			otpcodes.Email(email),
			otpcodes.ExparationTimeGT(time.Now()),
		).
		Order((ent.Desc("exparation_time"))).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, models.ErrOTPNotFound
		}
		return nil, err
	}
	return otpToModel(otp), err
}

func otpToModel(otp *ent.OtpCodes) *models.OTP {
	return &models.OTP{
		Code:           otp.Code,
		NextSendTime:   otp.NextSendTime,
		ExparationTime: otp.ExparationTime,
	}
}
