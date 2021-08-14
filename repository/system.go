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
	c := GetAllColumns()
	dbinfo := DBInfo{
		TableCount: len(t),
		DbName:     helpers.Cfg.DbName,
		Tables:     t,
		Columns:    c,
	}
	return dbinfo
}

func GetColumns(tableName string) []ShowColumns {
	var result []ShowColumns
	helpers.DB().Raw("SELECT m.name as Name, p.name as ColumnName, p.type as ColumnType FROM sqlite_master m left outer join pragma_table_info((m.name)) p on m.name <> p.name WHERE m.type ='table' AND m.name = ?", tableName).Scan(&result)
	return result
}

func GetAllColumns() []ShowColumns {
	var result []ShowColumns
	helpers.DB().Raw("SELECT p.name as Name, p.name as ColumnName, p.type as ColumnType FROM sqlite_master m left outer join pragma_table_info((m.name)) p on m.name <> p.name WHERE m.type ='table' AND m.name NOT LIKE 'sqlite_%'").Scan(&result)
	return result
}

func GetTables() []ShowTables {
	var result []ShowTables
	helpers.DB().Raw("SELECT m.name, m.sql FROM sqlite_master m WHERE m.type ='table' AND m.name NOT LIKE 'sqlite_%'").Scan(&result)
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

type ShowColumns struct {
	Name string
	ColumnName string
	ColumnType string
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
	Columns    []ShowColumns `json:"columns"`
}
