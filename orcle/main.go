package main

import (
	"fmt"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-oci8"

	"github.com/go-xorm/core"
)

type SYS_CONFIG struct {
	CFG_ID     string `xorm:"not null pk default '' VARCHAR(50) " `
	CFG_NAME   string `xorm:"default '' VARCHAR(100) "`
	CFG_DATA   string `xorm:"default '' VARCHAR(50) "`
	CFG_TYPE   string `xorm:"default '' VARCHAR(50) "`
	CFG_STATUS int    `xorm:"default 1 INT(2) "`
}


func main() {
	Engine, err := xorm.NewEngine("oci8", "root/123@192.168.0.120:1521/ORCL")
	if err != nil {
		fmt.Println("新建引擎", err)
		return
	}


	Engine.SetTableMapper(core.SameMapper{})
	Engine.SetColumnMapper(core.SameMapper{})

	Engine.ShowSQL(true)
	Engine.SetMaxIdleConns(5)

	 config  := new(SYS_CONFIG)
	isTrue, err := Engine.Get(config)
	if !isTrue || err != nil {
		fmt.Println("this is get error : ", err)
		return
	}
	fmt.Println(config)
}
