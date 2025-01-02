package task

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var db *gorm.DB

func Init(dsn string) (*gorm.DB, error) {
    var err error
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}