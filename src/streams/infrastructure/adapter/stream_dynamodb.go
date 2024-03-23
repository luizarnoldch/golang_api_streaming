package adapter

import (
	"context"
	"log"
	"main/src/streams/domain/model"
	"main/src/streams/domain/repository"
	appError "main/utils/error"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
)

type StreamDynamoDBRepository struct {
	client *dynamodb.Client
	ctx    context.Context
	table  string
}

func NewStreamDynamoDBRepository(ctx context.Context, client *dynamodb.Client, table string) repository.StreamRepository {
	return &StreamDynamoDBRepository{
		client: client,
		ctx:    ctx,
		table:  table,
	}
}

func (repo *StreamDynamoDBRepository) UpdateStreamById(stream_id string, stream *model.Stream) (*model.Stream, *appError.Error) {
	updateBuilder := expression.Set(
		expression.Name("name"),expression.Value(stream.Name)).Set(
		expression.Name("cost"), expression.Value(stream.Cost)).Set(
		expression.Name("start_date"), expression.Value(stream.StartDate)).Set(
		expression.Name("end_date"), expression.Value(stream.StartDate))
	expr, err := expression.NewBuilder().WithUpdate(updateBuilder).Build()
	if err != nil {
		log.Printf("Error building expression for UpdateStreamById, %s", err.Error())
		return nil, appError.NewUnexpectedError(err.Error())
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(repo.table),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: stream_id},
		},
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ReturnValues:              types.ReturnValueUpdatedNew,
	}

	output, err := repo.client.UpdateItem(repo.ctx, input)
	if err != nil {
		log.Printf("Error updating stream UpdateStreamById, %s", err.Error())
		return nil, appError.NewUnexpectedError(err.Error())
	}
	var updatedStream model.Stream
	if err := attributevalue.UnmarshalMap(output.Attributes, &updatedStream); err != nil {
		log.Printf("Error unmarshaling UpdateConnectInstancesItem, %s", err.Error())
		return nil, appError.NewUnexpectedError(err.Error())
	}
	return &updatedStream, nil
}

func (repo *StreamDynamoDBRepository) GetStreamById(stream_id string) (*model.Stream, *appError.Error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(repo.table),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: stream_id},
		},
	}
	result, err := repo.client.GetItem(repo.ctx, input)
    if err != nil {
        log.Printf("Error in GetItem: %s", err.Error())
        return nil, appError.NewUnexpectedError(err.Error())
    }
	if result.Item == nil {
		log.Printf("No item found with on GetStreamById, %s", stream_id)
		return nil, appError.NewUnexpectedError("No item found with on GetStreamById")
	}
	var stream model.Stream
	if err := attributevalue.UnmarshalMap(result.Item, &stream); err != nil {
		log.Printf("Error unmarshaling result: %s", err.Error())
		return nil, appError.NewUnexpectedError(err.Error())
	}
	return &stream, nil
}

func (repo *StreamDynamoDBRepository) GetAllStream() ([]model.Stream, *appError.Error) {
	var response []model.Stream
	input := &dynamodb.ScanInput{
		TableName: aws.String(repo.table),
	}
	output, err := repo.client.Scan(repo.ctx, input)
	if err != nil {
        log.Printf("Scan error in GetAllStream: %s", err.Error())
        return nil, appError.NewUnexpectedError(err.Error())
    }
	for _, item := range output.Items {
		var stream model.Stream
		err := attributevalue.UnmarshalMap(item, &stream)
		if err != nil {
			log.Printf("Error while unmarshaling item from GetAllCompanys, %s", err.Error())
			return nil, appError.NewUnexpectedError(err.Error())
		}
		response = append(response, stream)
	}
	return response, nil
}

func (repo *StreamDynamoDBRepository) CreateStream(stream *model.Stream) (*model.Stream, *appError.Error) {
	marshalStream, err := attributevalue.MarshalMap(stream)
	if err != nil {
		log.Printf("Error while marshalling CreateStream input, %s", err.Error())
		return nil, appError.NewUnexpectedError(err.Error())
	}
    input := &dynamodb.PutItemInput{
        Item:      marshalStream,
        TableName: aws.String(repo.table),
    }
	_, err = repo.client.PutItem(repo.ctx, input)
	if err != nil {
		log.Printf("Error while PutItem on CreateStream: %s", err.Error())
		return nil, appError.NewUnexpectedError(err.Error())
	}
	log.Printf("CreateStream done successfully")
	return stream, nil
}

func (repo *StreamDynamoDBRepository) DeleteStream(stream_id string) (*appError.Error) {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(repo.table),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: stream_id},
		},
	}

	_, err := repo.client.DeleteItem(repo.ctx, input)
	if err != nil {
        log.Printf("Error while DeleteStream, %s", err.Error())
        return appError.NewUnexpectedError(err.Error())
    }
	return nil
}