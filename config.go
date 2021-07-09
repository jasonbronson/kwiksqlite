package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"gorm.io/gorm"
)

// Cfg instance
var Cfg *Config = &Config{}

// Config struct
type Config struct {
	Port              string
	DisableSSL        bool
	GormDB            *gorm.DB
	GormDBReadReplica *gorm.DB
	CorsOption        *cors.Config
	HttpClient        *http.Client
	JwtConfig         *JWTConfig
}

func init() {

	Cfg.Port = os.Getenv("PORT")
	if Cfg.Port == "" {
		Cfg.Port = "1234"
	}
	Cfg.DisableSSL, _ = strconv.ParseBool(os.Getenv("DISABLE_SSL"))

	//Jwt setup
	jwt := JWTConfig{
		Secret:   os.Getenv("JWT_SECRET"),
		Issuer:   os.Getenv("JWT_ISSUER"),
		Audience: os.Getenv("JWT_AUDIENCE"),
	}
	Cfg.JwtConfig = &jwt

	//Setup cors
	Cors()

}

func Cors() {
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-Redirect-Host", "X-Paypal-Checkout", "X-odd-user-agent"},
		AllowCredentials: true,
	}

	Cfg.CorsOption = &corsConfig
}

type JWTConfig struct {
	Secret   string
	Issuer   string
	Audience string
}
