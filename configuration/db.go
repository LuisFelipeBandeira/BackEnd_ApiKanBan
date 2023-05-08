package configuration

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDb() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error to load .env file")
	}

	connectionString := os.Getenv("CONNECT_DB")

	DB, errConnectDB := sql.Open("mysql", connectionString)
	if errConnectDB != nil {
		return nil, errConnectDB
	}

	if errPingDb := DB.Ping(); errPingDb != nil {
		return nil, errPingDb
	}

	return DB, nil
}
