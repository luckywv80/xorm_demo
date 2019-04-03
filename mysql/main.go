package main

import (
	"fmt"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

type SysConfig struct {
	CfgId     string `xorm:"not null pk default '' VARCHAR(50)"`
	CfgName   string `xorm:"default '' VARCHAR(100)"`
	CfgData   string `xorm:"default '' VARCHAR(1024)"`
	CfgType   string `xorm:"default '' VARCHAR(50)"`
	CfgStatus int    `xorm:"default 1 INT(2)"`
}

func (s *SysConfig) TableName() string {
	return "sys_config"
}

func main() {

	Engine, err := xorm.NewEngine("mysql", "root:www.upsoft01.com@/antdbms?charset=utf8")
	if err != nil {
		fmt.Println("新建引擎", err)
		return
	}
	if err := Engine.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	Engine.SetTableMapper(core.SameMapper{})
	Engine.ShowSQL(true)
	Engine.SetMaxIdleConns(5)

	config := new(SysConfig)
	isTrue, err := Engine.Get(config)
	if !isTrue|| err != nil {
		fmt.Println("this is get error : ", err)
		return
	}
	fmt.Println(config)
}
