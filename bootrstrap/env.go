package bootrstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerHost string
	ServerPort string
	ServerUrl string
	SecretKey string
	Db         configDb
}

type configDb struct {
	DbHost string
	DbPort string
	DbPass string
	Dbname string
	DbUser string
}

type EnvStruct struct {
	ServerHost           string `mapstructure:"SERVER_HOST"`
	ServerPort           string `mapstructure:"SERVER_PORT"`
	ServerUrl           string `mapstructure:"SERVER_URL"`
	SecretKey           string `mapstructure:"SERVER_KEY"`
}

type dbConfig struct {
	DbHost string `mapstructure:"DB_HOST"`
	DbPort string `mapstructure:"DB_PORT"`
	DbPass string `mapstructure:"DB_PASS"`
	Dbname string `mapstructure:"DB_NAME"`
	DbUser string `mapstructure:"DB_USER"`
}

func NewEnv() *Env {
	env := EnvStruct{}
	dbConfig := dbConfig{}
	config := Env{}

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	err = viper.Unmarshal(&dbConfig)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	config.ServerHost = env.ServerHost
	config.ServerPort = env.ServerPort
	config.ServerUrl = env.ServerUrl
	config.SecretKey = env.SecretKey

	// db rajatracker pulsa
	db := &configDb{
		DbHost: dbConfig.DbHost,
		DbPort: dbConfig.DbPort,
		DbPass: dbConfig.DbPass,
		Dbname: dbConfig.Dbname,
		DbUser: dbConfig.DbUser,
	}
	config.Db = *db

	return &config
}
