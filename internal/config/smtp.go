package config

type SMTP struct {
	User              string `yaml:"user" validate:"required"`
	Password          string `yaml:"password" validate:"required"`
	Host              string `yaml:"host" validate:"required"`
	Port              string `yaml:"port" validate:"required"`
	OtpLength         int    `yaml:"otp_length" validate:"required"`
	NextSendPeriodSec int    `yaml:"next_send_period_sec" validate:"required"`
	ExparationTimeMin int    `yaml:"exparation_time_min" validate:"required"`
}
