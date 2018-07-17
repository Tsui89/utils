package adapter

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Adapter struct {
	driverName     string
	dataSourceName string
	dbName         string
	//createTables   []string
	engine *gorm.DB
}

func finalizer(a *Adapter) {
	a.engine.Close()
}

func NewEngine(driverName, dataSourceName, dbName string) *gorm.DB {
	a := &Adapter{}
	a.driverName = driverName
	a.dataSourceName = dataSourceName
	a.dbName = dbName
	//a.createTables = createTables

	// Open the DB, create it if not existed.
	a.open()
	// Call the destructor when the object is released.
	//runtime.SetFinalizer(a, finalizer)

	return a.engine
}

func (a *Adapter) createDatabase() error {
	var err error
	var engine *gorm.DB
	if a.driverName == "postgres" {
		engine, err = gorm.Open(a.driverName, a.dataSourceName+" dbname=postgres")
	} else {
		engine, err = gorm.Open(a.driverName, a.dataSourceName)
	}
	if err != nil {
		return err
	}
	defer engine.Close()

	if a.driverName == "postgres" {
		err = engine.Exec(fmt.Sprintf("CREATE DATABASE %s ;", a.dbName)).Error
	} else {
		err = engine.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARSET utf8 COLLATE utf8_general_ci;", a.dbName)).Error
	}
	return err
}

func (a *Adapter) open() {
	var err error
	var engine *gorm.DB

	if err = a.createDatabase(); err != nil {
		panic(err)
	}

	if a.driverName == "postgres" {
		engine, err = gorm.Open(a.driverName, a.dataSourceName+fmt.Sprintf(" dbname=%s", a.dbName))
	} else {
		engine, err = gorm.Open(a.driverName, a.dataSourceName+a.dbName+"?charset=utf8&parseTime=True&loc=Local")
	}
	if err != nil {
		panic(err)
	}

	a.engine = engine
}

//
//func (a *Adapter) createTable() {
//	for _, s := range a.createTables {
//		_, err := a.engine.Exec(s)
//		if err != nil {
//			panic(err)
//		}
//	}
//}
