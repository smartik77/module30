package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"
	"testing"
)

func TestCreateTask(t *testing.T) {
	type args struct {
		db         *pgxpool.Pool
		authorID   int
		assignedID int
		title      string
		content    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateTask(tt.args.db, tt.args.authorID, tt.args.assignedID, tt.args.title, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	type args struct {
		db     *pgxpool.Pool
		taskID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteTask(tt.args.db, tt.args.taskID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllTasks(t *testing.T) {
	type args struct {
		db *pgxpool.Pool
	}
	tests := []struct {
		name    string
		args    args
		want    []Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllTasks(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTasksByAuthor(t *testing.T) {
	type args struct {
		db       *pgxpool.Pool
		authorID int
	}
	tests := []struct {
		name    string
		args    args
		want    []Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTasksByAuthor(tt.args.db, tt.args.authorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTasksByAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTasksByAuthor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTasksByLabel(t *testing.T) {
	type args struct {
		db      *pgxpool.Pool
		labelID int
	}
	tests := []struct {
		name    string
		args    args
		want    []Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTasksByLabel(tt.args.db, tt.args.labelID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTasksByLabel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTasksByLabel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	type args struct {
		db      *pgxpool.Pool
		taskID  int
		title   string
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateTask(tt.args.db, tt.args.taskID, tt.args.title, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
