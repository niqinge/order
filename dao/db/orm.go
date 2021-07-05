package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"strings"
)

func NewOrm() *gorm.DB {
	un, pw, err := loadLocal()
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", un, pw, "test")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DryRun:  false,
		Plugins: nil,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	
	d, _ := db.DB()
	if err := d.Ping(); err != nil {
		panic(err)
	}
	return db
}

func loadLocal() (username, password string, err error) {
	userNameData, err := ioutil.ReadFile("/usr/local/.db/mysql.uname")
	if err != nil {
		return username, password, err
	}
	un := strings.TrimSpace(string(userNameData))
	username = un
	pwdData, err := ioutil.ReadFile("/usr/local/.db/mysql.pas")
	if err != nil {
		return username, password, err
	}
	pwd := strings.TrimSpace(string(pwdData))
	password = pwd
	return username, password, err
}
