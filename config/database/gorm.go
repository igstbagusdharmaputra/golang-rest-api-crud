package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type gormInstance struct {
	master *gorm.DB
}

func (g *gormInstance) Master() *gorm.DB {
	return g.master
}

type GormDatabase interface {
	Master() *gorm.DB
}

func InitGorm() GormDatabase {
	inst := new(gormInstance)
	gormConfig := &gorm.Config{}
	// username, password, host, port, database
	v := os.Getenv("MYSQL_PORT")
	port, _ := strconv.Atoi(v)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), port, os.Getenv("MYSQL_DBNAME"))

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: connection,
	}), gormConfig)

	if err != nil {
		fmt.Printf("cannot connection database : %s ", err)
		panic("connection error")
	}
	fmt.Println(connection)
	fmt.Println("DB Connection OK!")

	inst.master = db
	return inst
}
