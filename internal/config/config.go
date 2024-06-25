package config

import "github.com/team-inu/inu-backyard/infrastructure/database"

type SessionConfig struct {
	MaxAge     int
	Secret     string
	Prefix     string
	CookieName string
}

type AuthConfig struct {
	Session   SessionConfig
	Turnstile TurnstileConfig
}

type CorsConfig struct {
	AllowOrigins []string
}

type TurnstileConfig struct {
	SecretKey string
}

type ClientConfig struct {
	Auth AuthConfig
	Cors CorsConfig
}

type FiberServerConfig struct {
	Database database.GormConfig
	Client   ClientConfig
}
