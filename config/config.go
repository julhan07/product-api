package config

import "os"

type App struct {
	Env     string
	Host    string
	Port    string
	Version string
}

type DBConfiguration struct {
	DbName   string
	Host     string
	User     string
	Password string
	Port     string
}

type RedisConfiguration struct {
	Host     string
	Password string
	Port     string
}

type Configuration struct {
	App                App
	PsqSource          DBConfiguration
	PsqDestination     DBConfiguration
	RedisConfiguration RedisConfiguration
}

func GetConfiguration() Configuration {

	return Configuration{
		App: App{
			Env:     os.Getenv("ENV"),
			Host:    os.Getenv("APP_HOST"),
			Port:    os.Getenv("APP_PORT"),
			Version: os.Getenv("APP_VERSION"),
		},
		PsqSource: DBConfiguration{
			DbName:   os.Getenv("DB_SOURCE_DB_NAME"),
			Host:     os.Getenv("DB_SOURCE_HOST"),
			User:     os.Getenv("DB_SOURCE_USER"),
			Password: os.Getenv("DB_SOURCE_PASSWORD"),
			Port:     os.Getenv("DB_SOURCE_PORT"),
		},
		PsqDestination: DBConfiguration{
			DbName:   os.Getenv("DB_DESTINATION_DB_NAME"),
			Host:     os.Getenv("DB_DESTINATION_HOST"),
			User:     os.Getenv("DB_DESTINATION_USER"),
			Password: os.Getenv("DB_DESTINATION_PASSWORD"),
			Port:     os.Getenv("DB_DESTINATION_PORT"),
		},
		RedisConfiguration: RedisConfiguration{
			Host:     os.Getenv("REDIS_HOST"),
			Password: os.Getenv("REDIS_PASSWORD"),
			Port:     os.Getenv("REDIS_PORT"),
		},
	}
}
