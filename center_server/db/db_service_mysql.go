package db
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

import (
	"gos_server/config"
)

func DBService() *gorm.DB {

	//gdb, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	gdb, err := gorm.Open("mysql", "root:@(127.0.0.1:3306)/test111?charset=utf8&parseTime=True&loc=Local")
	//defer gdb.Close()
	if err != nil {
		config.LogCenterServer.Errorln("cannot connect to", config.DBInner, err)
		os.Exit(-1)
	}
	gdb.DB().SetMaxIdleConns(10)
	gdb.DB().SetMaxOpenConns(100)
	config.LogCenterServer.Infoln("DBService connect to ", config.DBInner," successfully!" )
	return gdb
}
