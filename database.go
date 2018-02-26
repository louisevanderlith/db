package db

import (
	"log"

	"github.com/astaxie/beego/orm"
)

// SyncDatabase will attempt to initialize tables
func SyncDatabase(dbSource string) {
	name := "default"
	driverName := "postgres"
	err := orm.RegisterDataBase(name, driverName, dbSource)

	if err != nil {
		log.Println("Please ensure that you have created your Database.")
	} else {
		orm.RunSyncdb(name, false, false)
	}
}
