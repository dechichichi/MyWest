package config

import (
	"gopkg,in/ini.v1"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	host           string
	port           string
	user           string
	password       string
	database       string
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
}
