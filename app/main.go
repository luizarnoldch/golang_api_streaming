package app

import (
	"context"
	"fmt"
	"main/config"

	stream_service "main/src/streams/application/service"
	stream_adapter "main/src/streams/infrastructure/adapter"

	user_service "main/src/users/application/service"
	user_adapter "main/src/users/infrastructure/adapter"

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

	stream_repository := stream_adapter.NewStreamDynamoDBRepository(
		ctx,
		localDynamoClient,
		config.MICRO.DB.STREAM_DYNAMODB.TABLE_NAME,
	)
	stream_application := stream_service.NewStreamDynamoDBService(stream_repository)
	stream_controller := StreamCotroller{stream_application}

	user_repository := user_adapter.NewUserDynamoDBRepository(
		ctx,
		localDynamoClient,
		config.MICRO.DB.USER_DYNAMODB.TABLE_NAME,
	)
	user_application := user_service.NewUserService(user_repository)
	user_controller := UserCotroller{user_application}

	app := fiber.New()

	app.Get("/", GreetHandler)

	api_stream := app.Group("/streams")
	api_stream.Post("/", stream_controller.CreateStream)
	api_stream.Get("/", stream_controller.GetAllStream)
	api_stream.Get("/:stream_id", stream_controller.GetStreamByID)
	api_stream.Put("/:stream_id", stream_controller.UpdateStream)
	api_stream.Delete("/:stream_id", stream_controller.DeleteStream)

	api_user := app.Group("/users")
	api_user.Post("/", user_controller.CreateUser)
	api_user.Get("/", user_controller.GetAllUser)
	api_user.Get("/:user_id", user_controller.GetUserByID)
	api_user.Put("/:user_id", user_controller.UpdateUserByID)
	api_user.Delete("/:user_id", user_controller.DeleteUser)

	URL_API := fmt.Sprint(config.MICRO.API.HOST,":",config.MICRO.API.PORT)
	app.Listen(URL_API)
}
