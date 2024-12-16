package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB инициализирует подключение к базе данных
func InitDB(dsn string) {
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Ошибка при проверке подключения:", err)
	}
	log.Println("Успешно подключено к базе данных!")
}
