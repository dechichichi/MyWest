package config

import (
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
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	LoadDb(file)
	//dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local", Dbuser, Dbpassword, Dbdatabase)
	dsn := "root:Ly05985481282@/todolistuser?charset=utf8mb4&parseTime=True&loc=Local"
	//localhost:Ly05985481282@/todolist?charset=utf8mb4&parseTime=True&loc=Local
	model.UserInit(dsn)
	dsn = "root:Ly05985481282@/todolistdata?charset=utf8mb4&parseTime=True&loc=Local"
	model.DataInit(dsn)
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
