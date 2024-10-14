package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

func main() {
	// Подключаемся к базе данных
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:5273@localhost:5432/test")
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	// Запросим данные из таблицы users
	rows, err := conn.Query(context.Background(), "SELECT id, name FROM author")
	if err != nil {
		log.Fatalf("query failed: %v", err)
	}
	defer rows.Close()

	// Обрабатываем результат
	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		fmt.Printf("ID: %d, name: %s\n", id, name)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("error occurred during iteration: %v", err)
	}
}
