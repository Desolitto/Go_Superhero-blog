package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"superhero-blog/db"
	"superhero-blog/pkg/handlers"
	"superhero-blog/pkg/utils"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	dsn = "host=localhost user=postgres dbname=postgres password=postgres port=5432 sslmode=disable"
)

func main() {
	utils.LoadAdminCredentials("admin_credentials.txt")

	db.InitDB(dsn) // Инициализация базы данных

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomePage).Methods("GET", "POST")
	router.HandleFunc("/admin", handlers.AdminPanel).Methods("GET", "POST")
	router.HandleFunc("/login", handlers.LoginPage).Methods("GET", "POST")
	router.HandleFunc("/article/{id}", handlers.ArticlePage).Methods("GET")
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images/"))))

	// Создаем канал для получения сигналов
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Запуск сервера в отдельной горутине
	go func() {
		fmt.Println("Сервер запущен на порту 8888...")
		if err := http.ListenAndServe(":8888", router); err != nil {
			fmt.Printf("Ошибка при запуске сервера: %v\n", err)
		}
	}()

	// Ожидание сигнала
	<-sigs
	fmt.Println("Получен сигнал для остановки сервера...")

}
