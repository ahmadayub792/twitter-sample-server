package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbconfig := DBConfig{
		Host:     "0.0.0.0",
		Port:     5432,
		User:     "postgres",
		DBName:   "postgres",
		Password: "docker",
	}

	return &dbconfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.DBName,
		dbConfig.Password,
	)
}

func ConnectDB() error {
	var err error
	DB, err = gorm.Open("postgres", DbURL(BuildDBConfig()))
	if err != nil {
		return fmt.Errorf("db connection failed: %w", err)
	}
	return nil
}

func CloseDB() {
	DB.Close()
}