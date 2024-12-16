// File: utils/auth.go

package utils

import (
	"bufio"
	"os"
	"strings"
)

var (
	AdminUser   string
	AdminPass   string
	CreateQuery string
)

func LoadAdminCredentials(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic("Не удалось открыть файл с учетными данными: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) != 2 {
			continue // Пропустить некорректные строки
		}
		switch parts[0] {
		case "admin_login":
			AdminUser = strings.TrimSpace(parts[1])
		case "admin_password":
			AdminPass = strings.TrimSpace(parts[1])
		case "create_table": // Загружаем запрос на создание таблицы в переменную
			CreateQuery = strings.TrimSpace(parts[1]) // Сохраняем запрос в переменной
		}
	}

	if err := scanner.Err(); err != nil {
		panic("Ошибка при чтении файла: " + err.Error())
	}
}
