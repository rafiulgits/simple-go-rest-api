package config

import (
	"sync"
)

//DBConfig : Database configuration structure
type DBConfig struct {
	Server   string
	Port     int
	User     string
	Password string
	DbName   string
}

var dbConfig *DBConfig

func mapDbConfig() {
	dbConfig = &DBConfig{
		Server:   "localhost",
		Port:     1433,
		User:     "",
		Password: "",
		DbName:   "GoRestApi",
	}
}

//NewDBConfig returns the database configuration instance
func NewDBConfig() *DBConfig {
	var loadDBOnce sync.Once
	loadDBOnce.Do(mapDbConfig)
	return dbConfig
}
