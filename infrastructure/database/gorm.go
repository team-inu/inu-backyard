package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GormConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func NewGorm(config *GormConfig) (gormDB *gorm.DB, err error) {
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	dsn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database
	gormDB, err = gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
