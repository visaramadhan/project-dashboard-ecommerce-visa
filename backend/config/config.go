package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppDebug   bool
	DBMigrate  bool
	DBSeeding  bool
}

func LoadConfig() (Config, error) {
	localEnv := viper.New()
	localEnv.SetConfigType("dotenv")
	viper.SetConfigFile("../.env") //specify config

	//set defaul value
	viper.SetDefault("DBHost", "localhost")
	viper.SetDefault("DBPort", "5432")
	viper.SetDefault("DBUser", "user")
	viper.SetDefault("DBPassword", "password")
	viper.SetDefault("DBName", "database")
	viper.SetDefault("AppDebug", true)
	viper.SetDefault("DBMigrate", false)
	viper.SetDefault("DBSeeding", false)

	//read env variabel
	viper.AutomaticEnv()
	//read configuration file
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file: %s, using default values or environment variables", err)
	}

	//add value to config
	config := Config{
		DBHost:     viper.GetString("DBHost"),
		DBPort:     viper.GetString("DBPort"),
		DBUser:     viper.GetString("DBUser"),
		DBPassword: viper.GetString("DBPassword"),
		DBName:     viper.GetString("DBName"),
		AppDebug:   viper.GetBool("AppDebug"),
		DBMigrate:  viper.GetBool("DBMigrate"),
		DBSeeding:  viper.GetBool("DBSeeding"),
	}
	return config, nil
}
