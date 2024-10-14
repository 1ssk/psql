package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

func main() {
	// Подключаемся к базе данных
	conn, err := pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/databasename")
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	// Запросим данные из таблицы users
	rows, err := conn.Query(context.Background(), "SELECT id, username, email, created_at FROM users")
	if err != nil {
		log.Fatalf("query failed: %v", err)
	}
	defer rows.Close()

	// Обрабатываем результат
	for rows.Next() {
		var id int
		var username string
		var email string
		var createdAt string // Замените на time.Time, если хотите использовать объект времени

		err := rows.Scan(&id, &username, &email, &createdAt)
		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		fmt.Printf("ID: %d, Username: %s, Email: %s, Created At: %s\n", id, username, email, createdAt)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("error occurred during iteration: %v", err)
	}
}
