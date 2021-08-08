package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jasonbronson/kwiksqlite-admin/helpers"
	"github.com/jasonbronson/kwiksqlite-admin/repository"
	"gorm.io/gorm"
)

func ConnectDB(g *gin.Context) {
	//helpers.Cfg.DbName = g.GetHeader("database")
	helpers.Cfg.DbName = g.Query("db")
	if !repository.CheckDBExists(helpers.Cfg.DbName) {
		g.JSON(500, gin.H{"error": "database file does not exist"})
		return
	}
	//log.Println(helpers.Cfg.DbName)
	d := repository.GetDatabaseInfo()
	success := false
	log.Println(d.TableCount)
	if d.TableCount > 0 {
		success = true
	}
	g.JSON(200, success)
}

func GetDatabaseInfo(g *gin.Context) {
	//TODO fix issue with missing dbstats
	d := repository.GetDatabaseInfo()
	g.JSON(200, d)
}

func GetTables(g *gin.Context) {
	var result []ShowTables
	helpers.DB().Raw("SELECT name, sql FROM sqlite_master WHERE type ='table' AND name NOT LIKE 'sqlite_%' order by name").Scan(&result)
	g.JSON(200, result)
}

func DropTable(g *gin.Context) {

	table := g.Query("table")
	//check if table exists
	if table == "" {
		g.JSON(500, gin.H{"error": "table name is required"})
		return
	}
	err := helpers.DB().Debug().Migrator().DropTable(table)
	if err != nil {
		log.Println(err)
		g.JSON(500, gin.H{"error": err.Error()})
		return
	}

	g.JSON(200, "Success")
}

func CreateTable(g *gin.Context) {
	table := g.Query("table")
	helpers.DB().Table(table).AutoMigrate(&NewTable{})
	g.JSON(200, "Success")
}

type ShowTables struct {
	Name string
	SQL  string
}

type NewTable struct {
	gorm.Model
}
