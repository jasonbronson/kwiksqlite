package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jasonbronson/kwiksqlite-admin/controller"
	"github.com/jasonbronson/kwiksqlite-admin/helpers"
)

//go:embed frontend/dist
var feserver embed.FS

func NewHTTPHandler(cfg *helpers.Config) http.Handler {
	router := gin.Default()

	//Middleware
	router.Use(cors.New(*cfg.CorsOption))
	//router.Use(controller.DBName())

	//Frontend serve vue ui
	router.Use(static.Serve("/", EmbedFolder(feserver, "frontend/dist")))
	router.NoRoute(func(c *gin.Context) {
		fmt.Printf("%s doesn't exist, redirect on /", c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	api := router.Group("/api")
	{
		api.GET("/tables", controller.GetTables)
		api.POST("/table/create", controller.CreateTable)
		api.DELETE("/table/drop/:tablename", controller.DropTable)
		api.GET("/databaseinfo", controller.GetDatabaseInfo)
		api.POST("/dbconnect", controller.ConnectDB)
		api.GET("/alive", func(g *gin.Context) {
			g.Writer.WriteHeader(http.StatusOK)
			g.Writer.Write([]byte(fmt.Sprintf("%d %s", http.StatusOK, http.StatusText(http.StatusOK))))
		})

	}

	return router
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	if err != nil {
		return false
	}
	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
