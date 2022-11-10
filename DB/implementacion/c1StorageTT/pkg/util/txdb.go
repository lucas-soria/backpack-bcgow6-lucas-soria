package util

import (
	"database/sql"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error: Loading .env")
	}
	var DSN = os.Getenv("DB_USER") + ":" + os.Getenv("DB_PSW") + "@/" + os.Getenv("DB")
	txdb.Register("txdb", "mysql", DSN)
}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("txdb", uuid.New().String())

	if err == nil {
		return db, db.Ping()
	}

	return db, err
}
