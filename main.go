//同一 三大数据库
package main

import (
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lunny/godbc"
	_ "github.com/mattn/go-oci8"

	"fmt"
)

type DbType struct {
	DbName   string
	DbSource string
}

type SYS_CONFIG struct {
	CFG_ID     string `xorm:"not null pk default '' VARCHAR(50) " `
	CFG_NAME   string `xorm:"default '' VARCHAR(100) "`
	CFG_DATA   string `xorm:"default '' VARCHAR(50) "`
	CFG_TYPE   string `xorm:"default '' VARCHAR(50) "`
	CFG_STATUS int    `xorm:"default 1 INT(2) "`
}


var dbType = []DbType{
	{DbName: "mysql", DbSource: "root:www.upsoft01.com@/antdbms?charset=utf8"},
	{DbName: "odbc", DbSource: "driver={SQL Server};Server=192.168.186.138;Database=antdbms;uid=sa;pwd=123;"}, //mssql
	{DbName: "oci8", DbSource: "root/123@192.168.0.120:1521/ORCL"},
}

var index  = 0  //mysql
//var index  = 0   //mssql
//var index  = 0  //oci8
func main() {


	fmt.Println("启动数据库数据库引擎 ",dbType[index].DbName)
	Engine, err := xorm.NewEngine(dbType[index].DbName, dbType[index].DbSource)

	if err != nil {
		fmt.Println("连接引擎出错 : ", err)
		return
	}

	Engine.SetTableMapper(core.SameMapper{})
	Engine.SetColumnMapper(core.SameMapper{})

	Engine.ShowSQL(true)
	Engine.SetMaxIdleConns(5)

	config :=new(SYS_CONFIG)
	if _,err=Engine.Get(config);err!=nil {
		fmt.Println(" get is error ",err)
	}
	fmt.Println(config)



}
