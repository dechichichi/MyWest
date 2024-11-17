package config

import (
	"strings"
	"todolist/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/ini.v1"
)

var (
	Dbhost     string
	Dbport     string
	Dbuser     string
	Dbpassword string
	Dbdatabase string
	AppMode    string
	HttpPort   string
)

func Init() {
	file, err := ini.Load("./config.ini")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	LoadDb(file)
	path := strings.Join([]string{Dbuser, ":", Dbpassword, "@", Dbhost, ":", Dbport, "/", Dbdatabase}, "")
	model.DbInit(path)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadDb(file *ini.File) {
	Dbhost = file.Section("mysql").Key("host").String()
	Dbport = file.Section("mysql").Key("port").String()
	Dbuser = file.Section("mysql").Key("user").String()
	Dbpassword = file.Section("mysql").Key("password").String()
	Dbdatabase = file.Section("mysql").Key("database").String()
}