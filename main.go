// 三大数据库
package main

import (
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lunny/godbc"
	_ "github.com/mattn/go-oci8"

	"log"
	"os"
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
	{DbName: "mysql", DbSource: "root:www.upsoft01.com@tcp(127.0.0.1:3306)/antdbms?charset=utf8"},
	{DbName: "odbc", DbSource: "driver={SQL Server};Server=192.168.186.138;Database=antdbms;uid=sa;pwd=123;"}, //mssql
	{DbName: "oci8", DbSource: "root/123@192.168.0.120:1521/ORCL"},
}

var index  = 0  //mysql
//var index  = 1   //mssql
//var index  = 2  //oci8
func main() {


	log.Println("启动数据库数据库引擎 ",dbType[index].DbName)
	Engine, err := xorm.NewEngine(dbType[index].DbName, dbType[index].DbSource)
	assert(err,"连接引擎出错")

	Engine.SetTableMapper(core.SameMapper{})
	Engine.SetColumnMapper(core.SameMapper{})

	Engine.ShowSQL(true)
	Engine.SetMaxIdleConns(5)

	config :=new(SYS_CONFIG)

	 _,err=Engine.Get(config)
	assert(err," get is error ")

	log.Println("通过结构体获取 : ",config,"\n\n")

	//获取单条
	config1  := make(map[string]string)
	//_,err = Engine.Table("SYS_CONFIG").Get(&config1)
	_,err = Engine.Table(new(SYS_CONFIG)).Get(&config1)
	assert(err,"通过数据错误报错 ")
	log.Println("通过条件获取 :",config1,"\n\n")

	//获取多条
	config2 :=make([]map[string]string,0)
	err=Engine.Table(new(SYS_CONFIG)).Find(&config2)
	assert(err,"获取多条报错 ")
	log.Println("获取多条",config2,"\n\n")

	//非struct的单个字段的数组
	cfgId :=make([]string,0)
	err = Engine.Table(new(SYS_CONFIG)).Cols("CFG_ID").Find(&cfgId)
	assert(err,"获取个别字段出错  ")
	log.Println("获取个别字段 ",cfgId,"\n\n")

	//非struct的多个字段的数组
	cfgIdList :=make([][]string,0)
	err = Engine.Table(new(SYS_CONFIG)).Cols("CFG_ID","CFG_TYPE").Find(&cfgIdList)
	assert(err,"获取多个字段出错  ")
	log.Println("获取多别字段 ",cfgIdList,"\n\n")

}


func assert(err error,msg string){
	if err != nil {
		log.Println(msg,err)
		os.Exit(0)
	}
}