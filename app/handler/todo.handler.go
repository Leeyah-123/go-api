package handler

import (
	"fiber/app/model"
	"fiber/app/service"
	"fiber/pkg/util"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var invalidID = "Invalid id"

func GetUserTodos(ctx *fiber.Ctx) error {
	// retrieving token meta data
	userID, err := retreiveUserId(ctx)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	todos, err := service.GetUserTodos(userID)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(todos)
}

func GetTodoById(ctx *fiber.Ctx) error {
	// retrieving token meta data
	userID, err := retreiveUserId(ctx)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	// retrieving todo id
	todoID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return service.ErrorResponse(fmt.Errorf(invalidID), ctx)
	}

	todo, err := service.GetTodoById(userID, todoID)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(todo)
}

func AddTodo(ctx *fiber.Ctx) error {
	var body model.Todo

	err := ctx.BodyParser(&body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	// validating request body
	errors := util.ValidateStruct(body)
	if len(errors) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// retrieving token meta data
	userID, err := retreiveUserId(ctx)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	todo, err := service.AddTodo(body, userID)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.Status(fiber.StatusCreated).JSON(todo)
}

func UpdateTodoById(ctx *fiber.Ctx) error {
	// retrieving token meta data
	userID, err := retreiveUserId(ctx)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	// retrieving todo id
	todoID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		service.ErrorResponse(fmt.Errorf(invalidID), ctx)
	}

	todo, err := service.GetTodoById(userID, todoID)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	updatedTodo, err := service.UpdateTodoById(todoID, !todo.Completed)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(updatedTodo)
}

func DeleteTodoById(ctx *fiber.Ctx) error {
	// retrieving token meta data
	userID, err := retreiveUserId(ctx)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	// retrieving todo id
	todoID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		service.ErrorResponse(fmt.Errorf(invalidID), ctx)
	}

	todo, err := service.GetTodoById(userID, todoID)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	err = service.DeleteTodoById(todo.ID)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func retreiveUserId(ctx *fiber.Ctx) (uuid.UUID, error) {
	// retrieving token meta data
	tokenData, err := util.ExtractTokenMetadata(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	return tokenData.ID, nil
}
