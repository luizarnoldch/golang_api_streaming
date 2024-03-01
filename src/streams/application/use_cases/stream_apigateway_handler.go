package usecases

import (
	"context"
	"log"
	"main/src/streams/application/service"
	"main/src/streams/domain/model"
	"main/src/streams/infrastructure/adapter"
	"main/src/streams/infrastructure/configuration"
	dynamodbUtils "main/utils/dynamodb"
	appError "main/utils/error"
)

func CreateStream(ctx context.Context, req *model.Stream) (*model.Stream, *appError.Error) {
	dynamoDBClient, err := dynamodbUtils.GetDynamoDBAWSClient(ctx)
	if err != nil {
		log.Printf("Failed to load DynamoDB from CreateStream API Function")
		return &model.Stream{}, appError.NewUnexpectedError("Failed to load DynamoDB from CreateStream API Function")
	}

	dynamoDBtable := configuration.GetDynamoDBStreamTable()

	StreamInfrastructure := adapter.NewStreamDynamoDBRepository(dynamoDBClient, ctx, dynamoDBtable)
	StreamService := service.NewStreamDynamoDBService(StreamInfrastructure)

	return StreamService.CreateStream(req)
}

func GetAllStream(ctx context.Context) ([]model.Stream, *appError.Error) {
	dynamoDBClient, err := dynamodbUtils.GetDynamoDBAWSClient(ctx)
	if err != nil {
		log.Printf("Failed to load DynamoDB from GetAllStream API Function")
		return []model.Stream{}, appError.NewUnexpectedError("Failed to load DynamoDB from GetAllStream API Function")
	}

	dynamoDBtable := configuration.GetDynamoDBStreamTable()

	StreamInfrastructure := adapter.NewStreamDynamoDBRepository(dynamoDBClient, ctx, dynamoDBtable)
	StreamService := service.NewStreamDynamoDBService(StreamInfrastructure)

	return StreamService.GetAllStream()
}

func GetItemById(ctx context.Context, stream_id string) (*model.Stream, *appError.Error) {
	dynamoDBClient, err := dynamodbUtils.GetDynamoDBAWSClient(ctx)
	if err != nil {
		log.Printf("Failed to load DynamoDB from GetAllStream API Function")
		return &model.Stream{}, appError.NewUnexpectedError("Failed to load DynamoDB from GetAllStream API Function")
	}

	dynamoDBtable := configuration.GetDynamoDBStreamTable()

	StreamInfrastructure := adapter.NewStreamDynamoDBRepository(dynamoDBClient, ctx, dynamoDBtable)
	StreamService := service.NewStreamDynamoDBService(StreamInfrastructure)

	return StreamService.GetStreamById(stream_id)
}
