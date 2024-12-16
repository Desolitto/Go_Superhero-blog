package handlers

import (
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"superhero-blog/db"
	"superhero-blog/models"
	"superhero-blog/pkg/utils"

	"github.com/gorilla/mux"
	"github.com/russross/blackfriday/v2"
)

var isAdminLoggedIn bool
var rateLimiter = make(chan struct{}, 100) // Ограничение на 100 запросов в секунду

func HomePage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		handleLogout(w, r)
		return
	}

	if !handleRateLimit(w) {
		return
	}

	page := getPageFromRequest(r)
	articles, hasNextPage := fetchArticles(page)

	funcMap := template.FuncMap{
		"safeHTML": func(html string) template.HTML {
			return template.HTML(html)
		},
	}

	tmpl, err := template.New("home_page.html").Funcs(funcMap).ParseFiles("template/home_page.html")
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{
		"Articles":        articles,
		"HasNextPage":     hasNextPage,
		"NextPage":        parsePage(page) + 1,
		"PrevPage":        parsePage(page) - 1,
		"IsAdminLoggedIn": isAdminLoggedIn, // Передаем состояние аутентификации
	})
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("logout") == "true" {
		isAdminLoggedIn = false
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func handleRateLimit(w http.ResponseWriter) bool {
	select {
	case rateLimiter <- struct{}{}:
		defer func() { <-rateLimiter }()
		return true
	default:
		http.Error(w, "429 Too Many Requests", http.StatusTooManyRequests)
		return false
	}
}

func getPageFromRequest(r *http.Request) string {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	return page
}

func fetchArticles(page string) ([]models.Article, bool) {
	const articlesPerPage = 3
	offset := (parsePage(page) - 1) * articlesPerPage

	rows, err := db.DB.Query("SELECT id, title, content FROM articles ORDER BY created_at DESC LIMIT $1 OFFSET $2", articlesPerPage, offset)
	if err != nil {
		return nil, false
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content); err != nil {
			return nil, false
		}
		formattedContent := FormatContent(string(blackfriday.Run([]byte(article.Content))))
		firstParagraph := extractFirstParagraph(formattedContent)
		truncatedContent := truncateContent(firstParagraph, 100)
		article.Content = truncatedContent
		articles = append(articles, article)
	}

	return articles, len(articles) == articlesPerPage
}

func ArticlePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var article models.Article
	err := db.DB.QueryRow("SELECT id, title, content FROM articles WHERE id = $1", id).Scan(&article.ID, &article.Title, &article.Content)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	// Преобразование HTML-контента
	article.Content = FormatContent(string(blackfriday.Run([]byte(article.Content))))

	funcMap := template.FuncMap{
		"safeHTML": func(html string) template.HTML {
			return template.HTML(html)
		},
	}
	t, err := template.New("article.html").Funcs(funcMap).ParseFiles("template/article.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, article)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}
func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Проверка учетных данных администратора
		if username == utils.AdminUser && password == utils.AdminPass {
			isAdminLoggedIn = true
			http.Redirect(w, r, "/admin?session=active", http.StatusSeeOther)
			return
		}

		// Проверка учетных данных обычного пользователя
		if username != "" && password != "" {

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		http.Error(w, "Неверные учетные данные", http.StatusUnauthorized)
		return
	}

	// Загрузка HTML-шаблона из файла
	tmplPath := "template/login.html"
	http.ServeFile(w, r, tmplPath)
}

func truncateContent(content string, maxLength int) string {
	if len(content) <= maxLength {
		return content
	}

	// Найдем последний пробел перед maxLength
	lastSpace := strings.LastIndex(content[:maxLength], " ")
	if lastSpace == -1 {
		return content[:maxLength] + "..."
	}

	return content[:lastSpace] + "..."
}

func extractFirstParagraph(content string) string {
	re := regexp.MustCompile(`<p>(.*?)</p>`)
	matches := re.FindStringSubmatch(content)
	if len(matches) > 1 {
		return matches[1]
	}
	return content // Если нет абзацев, возвращаем оригинальный контент
}

func parsePage(page string) int {
	p, err := strconv.Atoi(page)
	if err != nil {
		return 1 // Возвращаем первую страницу при ошибке парсинга
	}
	return p
}

func FormatContent(content string) string {
	// Разделяем контент на строки
	lines := strings.Split(content, "\n")
	var cleanedLines []string
	for _, line := range lines {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			cleanedLines = append(cleanedLines, "<p>"+trimmed+"</p>")
		}
	}
	return strings.Join(cleanedLines, "")
}
