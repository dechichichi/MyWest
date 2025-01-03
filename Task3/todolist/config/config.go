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

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig() //读取配置
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	config := &Config{}
	configDB := viper.Sub("mysql")
	err = configDB.Unmarshal(&config.DB)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	configService := viper.Sub("service")
	err = configService.Unmarshal(&config.Service)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return config, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DB.Dbuser, c.DB.Dbpassword, c.DB.Dbhost, c.DB.Dbport, c.DB.Dbdatabase)
}

func Init() (*Config, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}
