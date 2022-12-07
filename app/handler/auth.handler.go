package handler

import (
	"fiber/app/model"
	"fiber/app/service"
	"fiber/pkg/util"
	"fiber/postgres"

	"github.com/gofiber/fiber/v2"
)

func Profile(ctx *fiber.Ctx) error {
	// retrieving token meta data
	tokenData, err := util.ExtractTokenMetadata(ctx)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	user, err := service.GetUserById(tokenData.ID)

	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.JSON(user)
	// return ctx.JSON(fiber.Map{})
}

func Login(ctx *fiber.Ctx) error {
	var body model.Login

	err := ctx.BodyParser(&body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	// // validating the user
	errors := util.ValidateStruct(body)
	if len(errors) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// retrieving the token by passing request body
	token, err := service.Login(body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}
	return ctx.JSON(fiber.Map{
		"token": token,
	})
}

func SignUp(ctx *fiber.Ctx) error {
	var body postgres.AddUserParams
	// parsing response body
	err := ctx.BodyParser(&body)
	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	// validating request body
	errors := util.ValidateStruct(model.User(body))
	if len(errors) != 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// retrieving the token by passing request body and registering user
	token, err := service.Register(body)

	if err != nil {
		return service.ErrorResponse(err, ctx)
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}
