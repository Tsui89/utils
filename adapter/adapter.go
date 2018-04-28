package adapter

import (
	"github.com/go-xorm/xorm"
	"github.com/lib/pq"
	"fmt"
)

type Adapter struct {
	driverName     string
	dataSourceName string
	dbName         string
	createTables   []string
	engine         *xorm.Engine
}

func finalizer(a *Adapter) {
	a.engine.Close()
}

func NewEngine(driverName, dataSourceName, dbName string, createTables []string) *xorm.Engine {
	a := &Adapter{}
	a.driverName = driverName
	a.dataSourceName = dataSourceName
	a.dbName = dbName
	a.createTables = createTables

	// Open the DB, create it if not existed.
	a.open()

	// Call the destructor when the object is released.
	//runtime.SetFinalizer(a, finalizer)

	return a.engine
}

func (a *Adapter) createDatabase() error {
	var err error
	var engine *xorm.Engine
	if a.driverName == "postgres" {
		engine, err = xorm.NewEngine(a.driverName, a.dataSourceName+" dbname=postgres")
	} else {
		engine, err = xorm.NewEngine(a.driverName, a.dataSourceName)
	}
	if err != nil {
		return err
	}
	defer engine.Close()

	if a.driverName == "postgres" {
		if _, err = engine.Exec(fmt.Sprintf("CREATE DATABASE %s ;"), a.dbName); err != nil {
			// 42P04 is	duplicate_database
			if err.(*pq.Error).Code == "42P04" {
				return nil
			}
		}
	} else {
		_, err = engine.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARSET utf8 COLLATE utf8_general_ci;", a.dbName))
	}
	return err
}

func (a *Adapter) open() {
	var err error
	var engine *xorm.Engine

	if err = a.createDatabase(); err != nil {
		panic(err)
	}

	if a.driverName == "postgres" {
		engine, err = xorm.NewEngine(a.driverName, a.dataSourceName+fmt.Sprintf(" dbname=%s", a.dbName))
	} else {
		engine, err = xorm.NewEngine(a.driverName, a.dataSourceName+a.dbName)
	}
	if err != nil {
		panic(err)
	}

	a.engine = engine
	err = a.engine.Ping()
	if err != nil {
		panic(err)
	}
	a.createTable()
}

func (a *Adapter) createTable() {
	for _, s := range a.createTables {
		_, err := a.engine.Exec(s)
		if err != nil {
			panic(err)
		}
	}
}
