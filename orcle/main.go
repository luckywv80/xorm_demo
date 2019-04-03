package main

import (
	"fmt"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-oci8"

)

type SysConfig struct {
	CFGID     string `xorm:"not null pk default '' VARCHAR(50) 'CFG_ID'" `
	CfgName   string `xorm:"default '' VARCHAR(100) 'CFG_NAME'"`
	CfgData   string `xorm:"default '' VARCHAR(1024) 'CFG_DATA'"`
	CfgType   string `xorm:"default '' VARCHAR(50) 'CFG_TYPE'"`
	CfgStatus int    `xorm:"default 1 INT(2) 'CFG_STATUS'"`
}

func (s *SysConfig) TableName() string {
	return "SYS_CONFIG"
}

func main() {
	Engine, err := xorm.NewEngine("oci8", "root/123@192.168.0.120:1521/ORCL")
	if err != nil {
		fmt.Println("新建引擎", err)
		return
	}


	//Engine.SetTableMapper(core.SameMapper{})
	Engine.ShowSQL(true)
	Engine.SetMaxIdleConns(5)

	 config  := new(SysConfig)
	bool, err := Engine.Get(config)
	if !bool || err != nil {
		fmt.Println("this is get error : ", err)
		return
	}
	fmt.Println(config)
}
