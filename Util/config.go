package Util

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	Port                 string        `mapstructure:"PORT"`
	SecurityKey          string        `mapstructure:"SECURITY_KEY"`
	AccessTime           time.Duration `mapstructure:"ACCESS_TIME"`
	AuthorizationPresent string        `mapstructure:"AUTH"`
}

func LoadConfig(path string) (config Config, readConfigErr error) {
	//set env file path
	viper.AddConfigPath(path)
	//set config fiile Name and type
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	//read values from the file
	viper.AutomaticEnv()
	readConfigErr = viper.ReadInConfig()
	if readConfigErr != nil {
		log.Fatalln("Unable to Read Env Variables: ", readConfigErr)
	}

	readConfigErr = viper.Unmarshal(&config)

	return
}
