package config

import (
    "fmt"
    "gopkg.in/ini.v1"
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
    file, err := ini.Load("./config/config.yml")
    if err != nil {
        return nil, err
    }

    config := &Config{}
    err = file.Section("mysql").MapTo(&config.DB)
    if err != nil {
        return nil, err
    }

    err = file.Section("service").MapTo(&config.Service)
    if err != nil {
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