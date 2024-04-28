package config

import (
	"os"
	"strconv"

	"gorm.io/gorm/logger"
)

type PrimaryDBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	TimeZone string
	LogLevel logger.LogLevel
}

type ReplicaDBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	TimeZone string
	LogLevel logger.LogLevel
}

func NewPrimaryDBConfig() (*PrimaryDBConfig, error) {
	p, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))

	if err != nil {
		return nil, err
	}

	lgl := logger.Silent
	if os.Getenv("APP_ENV") == "local " || os.Getenv("APP_ENV") == "development" {
		lgl = logger.Info
	}

	return &PrimaryDBConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     p,
		Database: os.Getenv("MYSQL_DATABASE"),
		TimeZone: os.Getenv("MYSQL_TIME_ZONE"),
		LogLevel: lgl,
	}, nil
}

func NewReplicaDBConfig() (*ReplicaDBConfig, error) {
	p, err := strconv.Atoi(os.Getenv("MYSQL_PORT_REPLICA"))
	if err != nil {
		return nil, err
	}

	lgl := logger.Silent
	if os.Getenv("APP_ENV") == "local " || os.Getenv("APP_ENV") == "development" {
		lgl = logger.Info
	}

	return &ReplicaDBConfig{
		User:     os.Getenv("MYSQL_USER_REPLICA"),
		Password: os.Getenv("MYSQL_PASSWORD_REPLICA"),
		Host:     os.Getenv("MYSQL_HOST_REPLICA"),
		Port:     p,
		Database: os.Getenv("MYSQL_DATABASE_REPLICA"),
		TimeZone: os.Getenv("MYSQL_TIME_ZONE"),
		LogLevel: lgl,
	}, nil
}
