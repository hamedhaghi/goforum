package config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	c *Config
)

const (
	MYSQL = "mysql"
)

type Config struct {
	Router     *chi.Mux
	Storage    *gorm.DB
	HttpServer string
	HttpPort   string
}

func GetConfig() *Config {
	if c != nil {
		return c
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	c := &Config{
		Router: chi.NewMux(),
		Storage: func() *gorm.DB {
			driver := os.Getenv("DB_DRIVER")
			if driver == MYSQL {
				dsn := fmt.Sprintf(
					"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
					os.Getenv("DB_USER"),
					os.Getenv("DB_PASSWORD"),
					os.Getenv("DB_HOST"),
					os.Getenv("DB_PORT"),
					os.Getenv("DB_NAME"),
				)
				gorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
				if err != nil {
					log.Fatal(err)
				}
				return gorm
			}
			return nil
		}(),
		HttpServer: "0.0.0.0",
		HttpPort:   "3000",
	}
	return c
}
