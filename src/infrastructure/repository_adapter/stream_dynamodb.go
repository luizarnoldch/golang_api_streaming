package repositoryadapter

import (
	"context"
	"log"
	"main/src/domain/model"
	dynamodbUtils "main/utils/dynamodb"
	appError "main/utils/error"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

)

type StreamDynamoDBRepository struct {
	client *dynamodb.Client
	ctx    context.Context
	table  string
}

func NewStreamDynamoDBRepository(client *dynamodb.Client, ctx context.Context, table string) *StreamDynamoDBRepository {
	return &StreamDynamoDBRepository{
		client: client,
		ctx:    ctx,
		table:  table,
	}
}

func (repo *StreamDynamoDBRepository) GetAllStream() ([]model.Stream, *appError.Error) {
	var response []model.Stream

	input := &dynamodb.ScanInput{
		TableName: aws.String(repo.table),
	}

	output, _ := repo.client.Scan(repo.ctx, input)

	for _, item := range output.Items {
		stream, _ := dynamodbUtils.UnmarshalStream(item)
		response = append(response, *stream)
	}

	return response, nil
}

func (repo *StreamDynamoDBRepository) CreateStream(stream *model.Stream) (*model.Stream, *appError.Error) {
	marshalStream, _ := dynamodbUtils.MarshalMapStream(stream)

    putInput := &dynamodb.PutItemInput{
        Item:      marshalStream,
        TableName: aws.String(repo.table),
    }
	
	repo.client.PutItem(repo.ctx, putInput)

	return stream, nil
}

func (repo *StreamDynamoDBRepository) GetStreamById(stream_id string) (*model.Stream, *appError.Error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(repo.table),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: stream_id},
		},
	}

	result, _ := repo.client.GetItem(repo.ctx, input)

	if result.Item == nil {
		log.Printf("GetStreamById: No item found with ID: %s", stream_id)
		return nil, appError.NewUnexpectedError("GetStreamById: No stream found with ID")
	}

	stream, _ := dynamodbUtils.UnmarshalStream(result.Item)
	return stream, nil
}