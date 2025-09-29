package database

import (
	"database/sql"
	"fmt"
	"log"

	"boilerplate-api-go/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ DB connection failed:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("❌ Cannot reach DB:", err)
	}

	log.Println("✅ Connected to MySQL")
	return db
}