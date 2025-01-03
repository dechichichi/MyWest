package config

import (
	"fmt"
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

var DSN string

func LoadConfig() {
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig() //读取配置
	if err != nil {
		log.Fatalln(err)
		return 
	}

	config := &Config{}
	configDB := viper.Sub("mysql")
	err = configDB.Unmarshal(&config.DB)
	if err != nil {
		log.Fatalln(err)
		return 
	}
	configService := viper.Sub("service")
	err = configService.Unmarshal(&config.Service)
	if err != nil {
		log.Fatalln(err)
		return 
	}
	DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB.Dbuser, config.DB.Dbpassword, config.DB.Dbhost, config.DB.Dbport, config.DB.Dbdatabase)
}

