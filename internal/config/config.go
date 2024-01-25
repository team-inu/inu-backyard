package config

import "github.com/team-inu/inu-backyard/infrastructure/database"

type SessionConfig struct {
	MaxAge     int
	Secret     string
	Prefix     string
	CookieName string
}

type AuthConfig struct {
	Session SessionConfig
}

type ClientConfig struct {
	Auth AuthConfig
}

type FiberServerConfig struct {
	Database database.GormConfig
	Client   ClientConfig
}
