package stream

import (
	"context"
	"log"
	appError "main/utils/error"
	"main/src/application/service"
	"main/src/domain/model"
	"main/src/infrastructure/configuration"
	"main/src/infrastructure/repository_adapter"
)

func CreateStream(ctx context.Context, req *model.Stream) (*model.Stream, *appError.Error) {
	dynamoDBClient, err := configuration.GetDynamoDBAWSClient(ctx)
	if err != nil {
		log.Printf("Failed to lead DynamoDB from CreateStream API Function")
		return &model.Stream{}, appError.NewUnexpectedError("Failed to lead DynamoDB from CreateStream API Function")
	}

	dynamoDBtable := configuration.GetDynamoDBStreamTable()

	StreamInfrastructure := repositoryadapter.NewStreamDynamoDBRepository(dynamoDBClient,ctx,dynamoDBtable)
	StreamService := service.NewStreamDynamoDBService(StreamInfrastructure)

	return StreamService.CreateStream(req)
}