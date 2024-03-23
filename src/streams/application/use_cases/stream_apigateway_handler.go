package usecases

import (
	"context"
	"log"
	"main/src/streams/application/service"
	"main/src/streams/domain/model"
	"main/src/streams/infrastructure/adapter"
	dynamodbUtils "main/utils/dynamodb"
	appError "main/utils/error"
)

type StreamUseCases struct {
	Ctx       context.Context
	TableName string
}

func (stream *StreamUseCases) UpdateStreamById(stream_id string, req *model.Stream) (*model.Stream, *appError.Error) {
	dynamoDBClient, err := dynamodbUtils.GetDynamoDBAWSClient(stream.Ctx)
	if err != nil {
		log.Printf("Failed to load DynamoDB from GetAllStream API Function")
		return nil, appError.NewUnexpectedError(err.Error())
	}

	StreamInfrastructure := adapter.NewStreamDynamoDBRepository(stream.Ctx, dynamoDBClient, stream.TableName)
	StreamService := service.NewStreamDynamoDBService(StreamInfrastructure)

	return StreamService.UpdateStreamById(stream_id, req)
}

func (stream *StreamUseCases) GetItemById(stream_id string) (*model.Stream, *appError.Error) {
	dynamoDBClient, err := dynamodbUtils.GetDynamoDBAWSClient(stream.Ctx)
	if err != nil {
		log.Printf("Failed to load DynamoDB from GetAllStream API Function")
		return nil, appError.NewUnexpectedError(err.Error())
	}

	StreamInfrastructure := adapter.NewStreamDynamoDBRepository(stream.Ctx, dynamoDBClient, stream.TableName)
	StreamService := service.NewStreamDynamoDBService(StreamInfrastructure)

	return StreamService.GetStreamById(stream_id)
}

func (stream *StreamUseCases) GetAllStream() ([]model.Stream, *appError.Error) {
	dynamoDBClient, err := dynamodbUtils.GetDynamoDBAWSClient(stream.Ctx)
	if err != nil {
		log.Printf("Failed to load DynamoDB from GetAllStream API Function")
		return nil, appError.NewUnexpectedError(err.Error())
	}

	StreamInfrastructure := adapter.NewStreamDynamoDBRepository(stream.Ctx, dynamoDBClient, stream.TableName)
	StreamService := service.NewStreamDynamoDBService(StreamInfrastructure)

	return StreamService.GetAllStream()
}

func (stream *StreamUseCases) CreateStream(req *model.Stream) (*model.Stream, *appError.Error) {
	dynamoDBClient, err := dynamodbUtils.GetDynamoDBAWSClient(stream.Ctx)
	if err != nil {
		log.Printf("Failed to load DynamoDB from CreateStream API Function")
		return nil, appError.NewUnexpectedError("Failed to load DynamoDB from CreateStream API Function")
	}

	StreamInfrastructure := adapter.NewStreamDynamoDBRepository(stream.Ctx, dynamoDBClient, stream.TableName)
	StreamService := service.NewStreamDynamoDBService(StreamInfrastructure)

	return StreamService.CreateStream(req)
}

func (stream *StreamUseCases) DeleteStream(stream_id string) *appError.Error {
	dynamoDBClient, err := dynamodbUtils.GetDynamoDBAWSClient(stream.Ctx)
	if err != nil {
		log.Printf("Failed to load DynamoDB from CreateStream API Function")
		return appError.NewUnexpectedError(err.Error())
	}

	StreamInfrastructure := adapter.NewStreamDynamoDBRepository(stream.Ctx, dynamoDBClient, stream.TableName)
	StreamService := service.NewStreamDynamoDBService(StreamInfrastructure)

	return StreamService.DeleteStream(stream_id)
}
