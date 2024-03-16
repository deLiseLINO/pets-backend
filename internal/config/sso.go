package config

type SSO struct {
	TokenLifespanHours int    `yaml:"token_lifespan_hours"`
	SecretKey          string `yaml:"sercet_key"`
}
