package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	c *Config
)

const (
	MYSQL  = "mysql"
	SQLITE = "sqlite"
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

	c = new(Config)
	c.Router = chi.NewRouter()
	storage, err := initDatabase()
	if err != nil {
		panic(err)
	}
	c.Storage = storage
	c.HttpServer = "0.0.0.0"
	c.HttpPort = "3000"
	return c
}

func initDatabase() (*gorm.DB, error) {
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
			return nil, err
		}
		return gorm, nil
	}
	if driver == SQLITE {
		gorm, err := gorm.Open(sqlite.Open(os.Getenv("DB_FILE")), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return gorm, nil
	}

	return nil, errors.New("unknown database driver")
}
