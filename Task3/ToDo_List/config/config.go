package config

import (
	"todo_list/model"

	"github.com/go-ini/ini"
)

var (
	AppMode     string
	HttpPort    string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
	Db          string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassword  string
	DbName      string
)

func Init() {
	file, err := ini.Load("./config.ini")
	if err != nil {
		panic(err)
	}
	LoadServe(file)
	LoadMysql(file)
	path := string.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True&loc=Local"})
	model.DataBase(path)
}

func LoadServe(file *ini.File) {
	AppMode := file.Section("service").Key("AppMode").String()
	HttpPort := file.Section("service").Key("HttpPort").String()
	RedisAddr := file.Section("redis").Key("RedisAddr").String()
	RedisPw := file.Section("redis").Key("RedisPw").String()
	RedisDbName := file.Section("redis").Key("RedisDbName").String()
}

func LoadMysql(file *ini.File) {
	Db := file.Section("mysql").Key("Db").String()
	DbHost := file.Section("mysql").Key("DbHost").String()
	DbPort := file.Section("mysql").Key("DbPort").String()
	DbUser := file.Section("mysql").Key("DbUser").String()
	DbPassword := file.Section("mysql").Key("DbPassword").String()
	DbName := file.Section("mysql").Key("DbName").String()
}
