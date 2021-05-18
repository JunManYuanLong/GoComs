package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"ict.com/project.v1/config"
	"io/ioutil"
)

type (
	Config struct {
		DbUri   string `json:"db_uri"`
		DbType  string `json:"db_type"`
		Address string `json:"address"`
	}
)

const (
	CONF         = "f"
	DEFAULT_CONF = "./dev.json"
	CONF_DES     = "conf path"
)

func main() {

	content, err := ioutil.ReadFile("test/dev.json")

	fmt.Println(content)
	confPath := flag.String(CONF, DEFAULT_CONF, CONF_DES)
	flag.Parse()
	c := Config{}
	fmt.Println(*confPath)
	_ = config.Load(*confPath, &c)

	db, err := gorm.Open("mysql", c.DbUri)
	if err != nil {
		fmt.Println(err)
		fmt.Println("mysql conntect err")
		return
	}
	db.SingularTable(true)
	db.LogMode(true)

}
