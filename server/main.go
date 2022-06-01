package main

import (
	"cpa_network/internal/config"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	var cfg *config.Config
	var db *gorm.DB
	var err error

	if cfg, err = config.ReadConfig("resources/config.json"); err != nil {
		panic("failed to read config file: " + err.Error())
	}
	if db, err = createDbConnection(cfg.Mysql); err != nil {
		panic("failed to create database connection: " + err.Error())
	}
	fmt.Println("successfully connected to " + db.Name() + " database\n")

	address := fmt.Sprintf("%s:%d", cfg.Server.Network.Host, cfg.Server.Network.Port)
	router := mux.NewRouter()
	if err = http.ListenAndServe(address, router); err != nil {
		log.Fatalf("failed to listen and serve: %s", err.Error())
		panic("failed to listen and serve: " + err.Error())
	}
}

func createDbConnection(cfg config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Login, cfg.Password, cfg.Network.Host, cfg.Network.Port, cfg.DBName)
	fmt.Println(dsn)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
