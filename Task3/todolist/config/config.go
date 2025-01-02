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

var Dsn string

func Init() {
    file, err := ini.Load("./config/config.ini")
    if err != nil {
        panic(err)
    }
    LoadServer(file)
    LoadDb(file)
    Dsn = Dbuser + ":" + Dbpassword + "@tcp(" + Dbhost + ":" + Dbport + ")/" + Dbdatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
    model.UserInit(Dsn)
}

func LoadServer(file *ini.File) {
    AppMode = file.Section("service").Key("AppMode").String()
    HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadDb(file *ini.File) {
    Dbhost = file.Section("mysql").Key("Dbhost").String()
    Dbport = file.Section("mysql").Key("Dbport").String()
    Dbuser = file.Section("mysql").Key("Dbuser").String()
    Dbpassword = file.Section("mysql").Key("Dbpassword").String()
    Dbdatabase = file.Section("mysql").Key("Dbdatabase").String()
}