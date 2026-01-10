package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *sql.DB

func setupDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Не удалось загрузить данные из .env файла")
	}

	dsn := os.Getenv("DB_DSN")

	if dsn == "" {
		log.Fatal("Переменная окружения DB_DSN не задана")
	}

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Успешное подключение к базе данных!")
}
