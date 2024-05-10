package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server       *Server
		DB           *DB
		Cloudinary   *Cloudinary
		GenerativeAI *GenerativeAI
	}

	Server struct {
		Port      int
		JWTSecret string
	}

	DB struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}

	Cloudinary struct {
		CloudName string
		ApiKey    string
		ApiSecret string
	}

	GenerativeAI struct {
		ApiKey string
	}
)

var (
	configInstance *Config
)

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&configInstance); err != nil {
		panic(err)
	}

	return configInstance
}
