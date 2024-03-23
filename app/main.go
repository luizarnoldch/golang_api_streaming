package app

import (
	"context"
	"fmt"
	"main/config"
	"main/src/streams/application/service"
	"main/src/streams/infrastructure/adapter"
	dynamodbUtils "main/utils/dynamodb"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func GreetHandler(c *fiber.Ctx) error {
	return c.SendString("Hello this is the API for Test")
}

func Start() {
	config, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatal("Environment error while loading .env %s", err)
	}
	ctx := context.TODO()
	localDynamoClient, err := dynamodbUtils.GetLocalDynamoDBClient(ctx)
	if err !=nil {
		log.Fatal("Error while getting local dynamodb client %s", err)
	}

	stream_repository := adapter.NewStreamDynamoDBRepository(
		ctx,
		localDynamoClient,
		config.MICRO.DB.STREAM_DYNAMODB.TABLE_NAME,
	)

	stream_application := service.NewStreamDynamoDBService(stream_repository)

	stream_controller := StreamCotroller{stream_application}

	app := fiber.New()

	app.Get("/", GreetHandler)

	api_stream := app.Group("/streams")
	api_stream.Post("/", stream_controller.CreateStream)
	api_stream.Get("/", stream_controller.GetAllStream)
	api_stream.Get("/:stream_id", stream_controller.GetStreamByID)
	api_stream.Put("/:stream_id", stream_controller.UpdateStream)
	api_stream.Delete("/:stream_id", stream_controller.DeleteStream)

	URL_API := fmt.Sprint(config.MICRO.API.HOST,":",config.MICRO.API.PORT)
	app.Listen(URL_API)
}
