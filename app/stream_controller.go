package app

import (
	"main/src/streams/application/service"
	"main/src/streams/domain/model"

	"github.com/gofiber/fiber/v2"
)

type StreamCotroller struct {
	service service.StreamService
}

func (controller *StreamCotroller) CreateStream(c *fiber.Ctx) error {
	var stream model.Stream
	if err := c.BodyParser(&stream); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	new_stream, err := controller.service.CreateStream(&stream)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err.ToString(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(new_stream)
}

func (controller *StreamCotroller) GetAllStream(c *fiber.Ctx) error {
	streams, err := controller.service.GetAllStream()
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err.ToString(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(streams)
}

func (controller *StreamCotroller) GetStreamByID(c *fiber.Ctx) error {
	streamID := c.Params("stream_id")
	stream, err := controller.service.GetStreamById(streamID)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err.ToString(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(stream)
}

func (controller *StreamCotroller) DeleteStream(c *fiber.Ctx) error {
	streamID := c.Params("stream_id")
	err := controller.service.DeleteStream(streamID)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err.ToString(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Stream deleted",
	})
}

func (controller *StreamCotroller) UpdateStream(c *fiber.Ctx) error {
	streamID := c.Params("stream_id")
	var stream model.Stream
	if err := c.BodyParser(&stream); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	update_stream, err := controller.service.UpdateStreamById(streamID, &stream)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err.ToString(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(update_stream)
}
