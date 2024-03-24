package app

import (
	"main/src/users/application/service"
	"main/src/users/domain/model"

	"github.com/gofiber/fiber/v2"
)

type UserCotroller struct {
	service service.UserService
}

func (controller *UserCotroller) CreateUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	new_stream, err := controller.service.CreateUser(&user)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err.ToString(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(new_stream)
}

func (controller *UserCotroller) GetAllUser(c *fiber.Ctx) error {
	streams, err := controller.service.GetAllUser()
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err.ToString(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(streams)
}

func (controller *UserCotroller) GetUserByID(c *fiber.Ctx) error {
	streamID := c.Params("user_id")
	stream, err := controller.service.GetUserByID(streamID)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err.ToString(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(stream)
}

func (controller *UserCotroller) DeleteUser(c *fiber.Ctx) error {
	streamID := c.Params("user_id")
	err := controller.service.DeleteUser(streamID)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err.ToString(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted",
	})
}

func (controller *UserCotroller) UpdateUserByID(c *fiber.Ctx) error {
	streamID := c.Params("user_id")
	var stream model.User
	if err := c.BodyParser(&stream); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	update_stream, err := controller.service.UpdateUserByID(streamID, &stream)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err.ToString(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(update_stream)
}
