package helpers

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DB() *gorm.DB {

	if Cfg.DbName == "" {
		log.Println("Database name not set")
		return nil
	}

	if Cfg.GormDB == nil && Cfg.DbName != "" {
		log.Println("Database is not connected to ", Cfg.DbName)
		db, _ := DatabaseConnect(Cfg.DbName)
		return db
	}
	return Cfg.GormDB
}

func DatabaseConnect(dbName string) (*gorm.DB, error) {

	var err error
	Cfg.GormDB, err = gorm.Open(sqlite.Open(fmt.Sprintf("file:%v", dbName)), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error: %v", err)
		return nil, err
	}

	db, err := Cfg.GormDB.DB()
	if err != nil {
		log.Fatalf("Error: %v", err)
		return nil, err
	}

	db.SetMaxOpenConns(Cfg.MaxConnections)
	db.SetMaxIdleConns(Cfg.MaxIdleConnections)
	if err = db.Ping(); err != nil {
		log.Fatal("Master Database cannot connect ", err)
	}
	log.Println("Success connecting to master database ", dbName)
	Cfg.DbName = dbName

	return Cfg.GormDB, nil
}
