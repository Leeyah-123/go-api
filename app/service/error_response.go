package service

import (
	"fiber/app/model"

	"github.com/gofiber/fiber/v2"
)

func SqlErrorIgnoreNotFound(err error) error {
	if err == nil {
		return nil
	}
	if err.Error() == "sql: no rows in result set" {
		return nil
	}
	return err
}

func ErrorResponse(err error, ctx *fiber.Ctx) error {
	if err.Error() == "sql: no rows in result set" {
		return ctx.Status(fiber.StatusNotFound).JSON(model.ErrorResponse{
			Message: "no record found",
		})
	}

	//if err.Error() == "ERROR: insert or update on table \"nok\" violates foreign key constraint \"nok_nok_mid_fkey\" (SQLSTATE 23503)" {
	//	return ctx.Status(fiber.StatusNotFound).JSON(model.ErrorResponse{
	//		Message: "Foreign key violated",
	//	})
	//}

	return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
		Message: err.Error(),
	})
}
