package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http" // пакет для поддержки HTTP протокола
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	response := fmt.Sprintf("Hello, %s!", name)
	fmt.Fprint(w, response)
}

func main() {
	// здесь ваш код

	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/api/user", handler)

	// Запускаем веб-сервер на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}

}
