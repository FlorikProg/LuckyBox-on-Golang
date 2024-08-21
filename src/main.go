package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Получен запрос на домашнюю страницу")
	tmpl, err := template.ParseFiles("C:\\Users\\Roman\\Desktop\\Web\\src\\templates\\index.html")
	if err != nil {
		log.Printf("Ошибка при парсинге шаблона: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Ошибка при выполнении шаблона: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}
}

func incrementWinHandler(w http.ResponseWriter, r *http.Request) {
	err := updateWinCounter(1)
	if err != nil {
		log.Printf("Ошибка при увеличении счетчика выигрышей: %v", err)
		http.Error(w, "Ошибка при увеличении счетчика", http.StatusInternalServerError)
		return
	}
	response := map[string]string{"status": "success"}
	json.NewEncoder(w).Encode(response)
}

func handleRequest() {

	fs := http.FileServer(http.Dir("./file"))
	http.Handle("/file/", http.StripPrefix("/file/", fs))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/increment-win", incrementWinHandler)

	log.Println("Сервер запущен на порту :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}

func main() {
	createDatabase()
	handleRequest()
}
