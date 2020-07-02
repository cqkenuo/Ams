package model

import (
	"Ams/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var Db *gorm.DB
var once sync.Once

func getConnStr(setting config.Config)string{
	return fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		setting.DbConf.DbUser,
		setting.DbConf.DbPwd,
		setting.DbConf.DbHost,
		setting.DbConf.DbName,
		setting.DbConf.DbCharSet)
}

func GetAppDB(setting config.Config) *gorm.DB {
	once.Do(func() {
		var err error
		Db,err = gorm.Open(setting.DbConf.Dbms,getConnStr(setting))
		if err != nil{
			panic(fmt.Sprintf("数据库连接错误:%s",err))
		}
	})
	return Db
}
