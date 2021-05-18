package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"ict.com/project.v1/config"
	"ict.com/project.v1/server"
)

const (
	CONF         = "f"
	DEFAULT_CONF = "project.v1/config/prod.json"
	CONF_DES     = "conf path"
)

type (
	Config struct {
		DbUri   string `json:"db_uri"`
		DbType  string `json:"db_type"`
		Address string `json:"address"`
	}
)

func main() {
	confPath := flag.String(CONF, DEFAULT_CONF, CONF_DES)
	flag.Parse()
	c := Config{}
	config.Load(*confPath, &c)

	db, err := gorm.Open("mysql", c.DbUri)
	if err != nil {
		fmt.Println(err)
		fmt.Println("mysql conntect err")
		return
	}
	db.SingularTable(true)
	db.LogMode(true)

	conf := &server.Config{
		Addr:  c.Address,
		Db:    *db,
		DbUri: c.DbUri,
	}
	s := server.NewServer(conf)
	s.Start()
}
