package config

import (
	"crud-golang/config/database"
	"os"
	"strconv"
	"sync"
)

type config struct {
	dbGorm database.GormDatabase
	port   int
}

type Config interface {
	ServiceName() string
	DB() database.GormDatabase
	Port() int
	ENV() string
}

func NewConfig() Config {
	cfg := new(config)
	cfg.connectDB()
	return cfg
}
func (c *config) ServiceName() string {
	return os.Getenv("SERVICE_NAME")
}

func (c *config) connectDB() {
	var loadonce sync.Once
	loadonce.Do(func() {
		c.dbGorm = database.InitGorm()
	})
}
func (c *config) Port() int {
	v := os.Getenv("PORT")
	c.port, _ = strconv.Atoi(v)
	return c.port
}
func (c *config) ENV() string {
	return os.Getenv(`ENVIRONTMENT`)
}
func (c *config) DB() database.GormDatabase {
	return c.dbGorm
}
