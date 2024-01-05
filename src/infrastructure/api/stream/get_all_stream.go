package stream

import (
	"context"
	"log"
	"main/src/application/service"
	"main/src/domain/model"
	"main/src/infrastructure/configuration"
	"main/src/infrastructure/repository_adapter"
)

func GetAllStream(ctx context.Context) ([]model.Stream, error) {
	dynamoDBClient, err := configuration.GetDynamoDBAWSClient(ctx)
	if err != nil {
		log.Printf("Failed to lead DynamoDB from GetAllStream API Function")
		return []model.Stream{}, err
	}

	dynamoDBtable := configuration.GetDynamoDBStreamTable()

	StreamInfrastructure := repositoryadapter.NewStreamDynamoDBRepository(dynamoDBClient,ctx,dynamoDBtable)
	StreamService := service.NewStreamDynamoDBService(StreamInfrastructure)

	return StreamService.GetAllStream()
}