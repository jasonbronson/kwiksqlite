package helpers

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// Cfg instance
var Cfg *Config = &Config{}

// Config struct
type Config struct {
	Port               string
	DisableSSL         bool
	GormDB             *gorm.DB
	DbName             string
	CorsOption         *cors.Config
	HttpClient         *http.Client
	JwtConfig          *JWTConfig
	MaxConnections     int
	MaxIdleConnections int
}

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file please create one first")
	}

	Cfg.Port = os.Getenv("PORT")
	//default to 1234
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
	Cfg.MaxConnections = 2
	Cfg.MaxIdleConnections = 2

	//Setup cors
	Cors()

}

func Cors() {
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "Database"},
		AllowCredentials: true,
	}

	Cfg.CorsOption = &corsConfig
}

type JWTConfig struct {
	Secret   string
	Issuer   string
	Audience string
}
