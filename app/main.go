package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"ict.com/app/config"
	server2 "ict.com/app/server"
)

const (
	CONF        = "f"
	DefaultConf = "app/config/dev.json"
	ConfDes     = "conf path"
)

type (
	Config struct {
		DbUri   string `json:"db_uri"`
		DbType  string `json:"db_type"`
		Address string `json:"address"`
	}
)

func main() {
	confPath := flag.String(CONF, DefaultConf, ConfDes)
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

	conf := &server2.Config{
		Addr:  c.Address,
		Db:    *db,
		DbUri: c.DbUri,
	}
	s := server2.NewServer(conf)
	s.Start()
}
