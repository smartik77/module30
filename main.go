package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"module30/pkg/storage"
)

/*
	hostname 127.0.0.1
	port 5432
	db mydatabase
	user testdb
	password testpass
	connStr := "postgres://testdb:testpass@127.0.0.1:5432/mydatabase"
*/

func main() {
	pwd := "testpass"
	connStr := "postgres://testdb:" + pwd + "@127.0.0.1:5432/mydatabase"

	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer db.Close()
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}
	fmt.Println("Successfully connected to the database!")

	// создание одной задачи
	_ = storage.CreateTask(db, 3, 3, "WW TITLE", "WW CONTENT")

	// получение всех задач
	tasks, _ := storage.GetAllTasks(db)
	fmt.Println(tasks)

	// получение задач по автору
	tasksByAuthor, _ := storage.GetTasksByAuthor(db, 4)
	fmt.Println(tasksByAuthor)

	// получение задач по лабелу
	tasksByLabel, _ := storage.GetTasksByLabel(db, 3)
	fmt.Println(tasksByLabel)

	// изменение задачи по id
	_ = storage.UpdateTask(db, 3, "UPDATE TITLE", "UPDATE CONTENT")

	// удаление задачи по id
	_ = storage.DeleteTask(db, 9)
}
