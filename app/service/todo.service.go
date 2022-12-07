package service

import (
	"context"
	"fiber/app/model"
	"fiber/platform/db"
	"fiber/postgres"

	"github.com/google/uuid"
)

func GetUserTodos(userID uuid.UUID) ([]postgres.Todo, error) {
	return db.Query().GetUserTodos(context.Background(), userID)
}

func GetTodoById(userID uuid.UUID, todoID uuid.UUID) (postgres.Todo, error) {
	return db.Query().GetTodoById(context.Background(), postgres.GetTodoByIdParams{
		ID:     todoID,
		UserID: userID,
	})
}

func AddTodo(todo model.Todo, userID uuid.UUID) (postgres.Todo, error) {
	return db.Query().AddTodo(context.Background(), postgres.AddTodoParams{
		UserID: userID,
		Title:  todo.Title,
	})
}

func UpdateTodoById(todoID uuid.UUID, status bool) (postgres.Todo, error) {
	return db.Query().UpdateTodoById(context.Background(), postgres.UpdateTodoByIdParams{
		ID:        todoID,
		Completed: status,
	})
}

func DeleteTodoById(todoID uuid.UUID) error {
	return db.Query().DeleteTodoById(context.Background(), todoID)
}
