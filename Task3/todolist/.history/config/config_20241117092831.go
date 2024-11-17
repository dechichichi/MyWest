package config

import (
	"gopkg.in/ini.v1"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Dbhost         string
	Dbport         string
	Dbuser         string
	Dbpassword     string
	Dbdatabase     string
	AppMode        string
	HttpPort       string
	redis_host     string
	redis_port     string
	redis_password string
	redisDbName    string
)

type Data struct {
	title      string
	state      bool
	content    string
	views      int
	status     int
	created_at string
	start_time string
	end_time   string
}

type User struct {
	ID       string
	Password string
	Data     Data
}

func Init() {
	file, err := ini.Load("config.ini")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	db, err := gorm.Open("mysql", "localhost:Ly05985481282@/todolist?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
	redis_host = file.Section("redis").Key("host").String()
	redis_port = file.Section("redis").Key("port").String()
	redis_password = file.Section("redis").Key("password").String()
	redisDbName = file.Section("redis").Key("dbName").String()

}

func LoadDb(file *ini.File) {
	Dbhost = file.Section("mysql").Key("host").String()
	Dbport = file.Section("mysql").Key("port").String()
	Dbuser = file.Section("mysql").Key("user").String()
	Dbpassword = file.Section("mysql").Key("password").String()
	Dbdatabase = file.Section("mysql").Key("database").String()
}
