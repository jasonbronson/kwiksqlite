package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/jasonbronson/kwiksqlite-admin/helpers"
	"gorm.io/gorm"
)

func CheckDBExists(dbName string) bool {
	if _, err := os.Stat(dbName); err == nil {
		return true
	} else {
		return false
	}
}

func GetDatabaseInfo() DBInfo {
	//TODO missing compiled stats https://github.com/mattn/go-sqlite3/issues/886
	t := GetTables()
	dbinfo := DBInfo{
		TableCount: len(t),
		DbName:     helpers.Cfg.DbName,
		Tables:     t,
	}
	return dbinfo
}

func GetTables() []ShowTables {
	var result []ShowTables
	helpers.DB().Raw("SELECT name, sql FROM sqlite_master WHERE type ='table' AND name NOT LIKE 'sqlite_%' order by name").Scan(&result)
	return result
}

func DropTable(tableName string) error {
	//check if table exists
	if tableName == "" {
		return fmt.Errorf("tableName is missing %s", tableName)
	}
	err := helpers.DB().Debug().Migrator().DropTable(tableName)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("drop Table error: %v", err)
	}
	return nil
}

func CreateTable(tableName string) error {
	err := helpers.DB().Table(tableName).AutoMigrate(&NewTable{})
	if err != nil {
		return err
	}
	return nil
}

type ShowTables struct {
	Name string
	SQL  string
}

type NewTable struct {
	gorm.Model
}

type DBInfo struct {
	TableCount int          `json:"table_count"`
	DbName     string       `json:"db_name"`
	Tables     []ShowTables `json:"tables"`
}
