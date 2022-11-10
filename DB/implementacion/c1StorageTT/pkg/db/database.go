package db

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func ConnectDatabase() (engine *gin.Engine, db *sql.DB) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: Loading .env")
	}
	configDB := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PSW"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DB"),
	}
	db, err = sql.Open("mysql", configDB.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	engine = gin.Default()
	return
}
