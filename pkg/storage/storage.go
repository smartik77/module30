package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Task struct {
	ID         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

func CreateTask(db *pgxpool.Pool, authorID, assignedID int, title, content string) error {
	query := `
        INSERT INTO tasks (author_id, assigned_id, title, content)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	var taskID int
	err := db.QueryRow(context.Background(), query, authorID, assignedID, title, content).Scan(&taskID)
	if err != nil {
		return fmt.Errorf("failed to create task: %v", err)
	}
	fmt.Printf("Task created with ID: %d\n", taskID)
	return nil
}

func GetAllTasks(db *pgxpool.Pool) ([]Task, error) {
	query := `
        SELECT id, opened, closed, author_id, assigned_id, title, content FROM tasks
    `
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %v", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Opened, &task.Closed, &task.AuthorID, &task.AssignedID, &task.Title, &task.Content)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %v", err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTasksByAuthor(db *pgxpool.Pool, authorID int) ([]Task, error) {
	query := `
		SELECT id, opened, closed, assigned_id, title, content FROM tasks
		WHERE author_id = $1
	`
	rows, err := db.Query(context.Background(), query, authorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks by author: %v", err)
	}

	var tasks []Task

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Opened, &task.Closed, &task.AssignedID, &task.Title, &task.Content)
		if err != nil {
			return nil, fmt.Errorf("failed to get tasks by author: %v", err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func GetTasksByLabel(db *pgxpool.Pool, labelID int) ([]Task, error) {
	query := `
		SELECT t.id, t.opened, t.closed, t.author_id, t.assigned_id, t.title, t.content
        FROM tasks t
        JOIN tasks_labels tl ON t.id = tl.task_id
        WHERE tl.label_id = $1
	`
	rows, err := db.Query(context.Background(), query, labelID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks by label: %v", err)
	}

	var tasks []Task

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Opened, &task.Closed, &task.AuthorID, &task.AssignedID, &task.Title, &task.Content)
		if err != nil {
			return nil, fmt.Errorf("failed to get tasks by label: %v", err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func UpdateTask(db *pgxpool.Pool, taskID int, title, content string) error {
	query := `
		UPDATE tasks
		SET title = $1, content = $2
		WHERE id = $3
	`
	_, err := db.Exec(context.Background(), query, title, content, taskID)
	if err != nil {
		return fmt.Errorf("failed to update task by id: %v", err)
	}

	fmt.Printf("Task with ID %d updated\n", taskID)
	return nil
}

func DeleteTask(db *pgxpool.Pool, taskID int) error {
	query := `
		DELETE FROM tasks WHERE id = $1
	`
	_, err := db.Exec(context.Background(), query, taskID)
	if err != nil {
		return fmt.Errorf("failed to dalete task by id: %v", err)
	}

	fmt.Printf("Task with ID %d deleted\n", taskID)
	return nil
}
