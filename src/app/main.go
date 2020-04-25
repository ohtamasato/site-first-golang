package main

import (
	_ "app/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func init() {
	orm.RegisterDriver(beego.AppConfig.String("driver"), orm.DRMySQL)
	orm.RegisterDataBase("default", beego.AppConfig.String("driver"), beego.AppConfig.String("sqlconn")+"?charset=utf8")
	err := orm.RunSyncdb("default", false, false)
	if err != nil {
		fmt.Println(err)
	}
}
