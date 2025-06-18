package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/rahulmjain24/go-server/config"
)

var DB *sql.DB

func Init() {
	dsn := config.GetEnv("DATABASE_URL", "")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in environment or .env")
	}

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	fmt.Println("Connected to PostgreSQL")
}
