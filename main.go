package main

import (
	"fmt"
	"runtime"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/TaXingTianJi/serverFramework/core"
	"github.com/TaXingTianJi/serverFramework/utils"

	_ "ssDemoServer/models"
	_ "ssDemoServer/msgs"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	core.ServerApp.Version("ssDemoServer")

	// orm
	core.SConfig.DBConf.User = "user"
	core.SConfig.DBConf.PW = "pw"
	//core.SConfig.DBConf.Addr = "localhost"
	//core.SConfig.DBConf.Port = 3306
	//core.SConfig.DBConf.DB = "testDb"
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		core.SConfig.DBConf.User, core.SConfig.DBConf.PW,
		core.SConfig.DBConf.Addr, core.SConfig.DBConf.Port,
		core.SConfig.DBConf.DB)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", str)
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)

	core.Run()
	//core.Run("127.0.0.1:60060")
	//core.Run("localhost")
	//core.Run(":60060")

	var wg utils.WaitGroupWrapper
	wg.Wrap(func() {
		serverRoom()
	})
	wg.Wait()
}

func serverRoom() {
	for {

	}
}

