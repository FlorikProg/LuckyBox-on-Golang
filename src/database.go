package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func createDatabase() {
	// Открытие соединения с базой данных
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		log.Fatalf("ошибка при открытии базы данных: %v", err)
	}
	defer db.Close()

	// Создание таблицы
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        win_counter INTEGER DEFAULT 0
    );`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("ошибка при создании таблицы: %v", err)
	}
	fmt.Println("таблица создана успешно!")
}

func updateWinCounter(id int) error {
	// Открытие соединения с базой данных
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		return fmt.Errorf("ошибка при открытии базы данных: %v", err)
	}
	defer db.Close()

	// Обновление счетчика выигрышей
	_, err = db.Exec("UPDATE users SET win_counter = win_counter + 1 WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении счетчика выигрышей: %v", err)
	}

	return nil
}
