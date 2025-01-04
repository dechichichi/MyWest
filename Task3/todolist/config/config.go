package config

import (
	"log"

	"github.com/spf13/viper"
)

type DBData struct {
	Dbhost     string
	Dbport     string
	Dbuser     string
	Dbpassword string
	Dbdatabase string
}

type ServiceData struct {
	AppMode  string
	HttpPort string
}

type Config struct {
	DB      DBData
	Service ServiceData
}

var ConfigData *Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig() //读取配置
	if err != nil {
		log.Fatalln(err)
		return
	}
	ConfigData = &Config{}
	configDB := viper.Sub("mysql")
	err = configDB.Unmarshal(&ConfigData.DB)
	if err != nil {
		log.Fatalln(err)
		return
	}
	configService := viper.Sub("service")
	err = configService.Unmarshal(&ConfigData.Service)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
