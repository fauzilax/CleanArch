package config

import (
	"log"

	"github.com/spf13/viper"
)

// deklarasikan Variable Global untuk diisikan token JWT KEY dari file local env
var (
	JWTKey = ""
)

type DBConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort int
	DBName string
	jwtKey string
}

func InitConfig() *DBConfig {
	return ReadEnv()
}

func ReadEnv() *DBConfig {
	app := DBConfig{}
	viper.SetConfigName("local")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("error read config : ", err.Error())
		return nil
	}
	err = viper.Unmarshal(&app)
	if err != nil {
		log.Println("error parse config : ", err.Error())
		return nil
	}

	JWTKey = app.jwtKey //isikan variable global JWTKey melalui fungsi ReadEnv yang diambil dari file Local.env

	return &app
}
