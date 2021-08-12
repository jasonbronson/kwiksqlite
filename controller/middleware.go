package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jasonbronson/kwiksqlite-admin/helpers"
)

func DBName() gin.HandlerFunc {
	return func(g *gin.Context) {
		dbName := g.Request.Header.Get("Database")
		if dbName != "" {
			log.Println("middleware setting database:", dbName)
			db, err := helpers.DatabaseConnect(dbName)
			if err != nil {
				log.Fatal(err)
			}
			helpers.Cfg.GormDB = db
			//helpers.SetCache("dbname", dbName, -1)
			return
		}

		//check cache
		// name, ok := helpers.GetCache("dbname")
		// if ok {
		// 	db, err := helpers.DatabaseConnect(fmt.Sprintf("%s", name))
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	helpers.Cfg.GormDB = db

		// }

	}
}
