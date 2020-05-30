/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/5/20 14:27
 */

package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

var (
	Engine *xorm.Engine
)

func init() {
	InitAppDb()
}

func InitAppDb() {
	args :=  "root:12345678@tcp(127.0.0.1:3306)/stations?charset=utf8"
	//args :=  "root:123456@tcp(127.0.0.1:3306)/stations?charset=utf8"
	Engine, _ = xorm.NewEngine("mysql", args)

	//设置日志
	Engine.ShowSQL(true)

	Engine.SetMaxIdleConns(10)
	Engine.SetConnMaxLifetime(time.Minute * 5)

	//保持心跳
	//go KeepHeart(Engine, "fapp_ads_info")
}