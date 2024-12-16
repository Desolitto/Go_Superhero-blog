package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"superhero-blog/db"
	"superhero-blog/pkg/utils"

	"github.com/russross/blackfriday/v2"
)

func AdminPanel(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Проверка на выход
		if r.FormValue("logout") == "true" {
			isAdminLoggedIn = false                       // Обновляем состояние аутентификации
			http.Redirect(w, r, "/", http.StatusSeeOther) // Перенаправляем на главную страницу
			return
		}

		title := r.FormValue("title")
		content := r.FormValue("content")

		// Преобразование содержимого Markdown в HTML
		htmlContent := blackfriday.Run([]byte(content))

		_, err := db.DB.Exec("INSERT INTO articles (title, content) VALUES ($1, $2)", title, htmlContent)
		if err != nil {
			http.Error(w, "Error saving article", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Проверка, вошел ли администратор
	if !checkAdminLoggedIn(w, r) {
		return
	}

	tmpl, err := template.ParseFiles("template/admin_panel.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Сбрасываем состояние аутентификации
	isAdminLoggedIn = false
	http.Redirect(w, r, "/", http.StatusSeeOther) // Перенаправляем на главную страницу
}

func CreateTables(db *sql.DB) {
	query := utils.CreateQuery
	if _, err := db.Exec(query); err != nil {
		panic("Ошибка при создании таблицы: " + err.Error())
	}
	fmt.Println("Таблица создана или уже существует.")
}

// Функция для проверки аутентификации администратора
func checkAdminLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	if !isAdminLoggedIn {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return false
	}
	return true
}
