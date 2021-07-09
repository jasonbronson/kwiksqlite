package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewHTTPHandler(cfg *Config) http.Handler {
	router := gin.Default()

	//Middleware
	router.Use(cors.New(*cfg.CorsOption))

	router.GET("/", func(g *gin.Context) {
		g.Writer.WriteHeader(http.StatusOK)
		g.Writer.Write([]byte(fmt.Sprintf("%d %s", http.StatusOK, http.StatusText(http.StatusOK))))
	})

	return router
}
