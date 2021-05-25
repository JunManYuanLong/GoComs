package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/urfave/cli/v2"
	"ict.com/project.v1/model"
	model2 "ict.com/public.v1/model"
	"os"
	"time"
)

const (
	DBURI = "root:root123@(ip:3306)/ep-go?charset=utf8&parseTime=True"
)

type (
	Manager struct {
		Db *gorm.DB
	}
)

func getDb() *gorm.DB {
	db, err := gorm.Open("mysql", DBURI)
	fmt.Println("db_uri is %s", DBURI)
	db.SingularTable(true)
	db.LogMode(true)
	if err != nil {
		fmt.Println("connect error")
	}
	return db
}

func createDb(c *cli.Context) error {
	db := getDb()
	fmt.Print(db)
	db.CreateTable(&model.Project{})
	return nil
}

func initData(c *cli.Context) error {
	db := getDb()
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err.Error())
	}
	p1 := &model.Project{EntityModel: model2.EntityModel{Status: 0, CreateTime: time.Now().In(local), UpdateTime: time.Now().In(local)}, Name: "EP", Description: "Devops一站式平台", Logo: "", Category: "inner", Level: 1, PM: 1, TD: 2, Background: "xxxxxx", Worth: "", Target: "让研发工程更具效能，让成本核算更方面", Milestone: "", Budget: 500000000000, IsShow: 1}
	db.SingularTable(true)
	db.LogMode(true)
	db.Create(p1)
	//db.Commit()
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "EP-GO"
	app.Usage = "manager file fo ep-go"
	app.Commands = cli.Commands{
		{
			Name:   "createdb",
			Usage:  "create database",
			Action: createDb,
		},
		{
			Name:   "initdata",
			Usage:  "init db data",
			Action: initData},
	}
	app.Run(os.Args)
}
