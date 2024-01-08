package repositoryadapter

import (
	"context"
	"log"
	"main/src/domain/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
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

func (repo *StreamDynamoDBRepository) GetAllStream() ([]model.Stream, error) {
	var response []model.Stream

	input := &dynamodb.ScanInput{
		TableName: aws.String(repo.table),
	}

	output, err := repo.client.Scan(repo.ctx, input)

	if err != nil {
		log.Printf("Error while scanning streams on GetAllStream with dynamodb stream_table: %v", err)
		return []model.Stream{}, nil
	}

	for _, item := range output.Items {
		var stream model.Stream
		err := attributevalue.UnmarshalMap(item, &stream)
		if err != nil {
			log.Printf("Error while marshaling stream by attributevalue: %v", err)
			return []model.Stream{}, nil
		}
		response = append(response, stream)
	}

	return response, nil
}

func (repo *StreamDynamoDBRepository) CreateStream(stream *model.Stream) (*model.Stream, error) {
	marshal_stream, err := attributevalue.MarshalMap(stream)

	if err != nil {
		log.Printf("Error while marshaling stream by attributevalue: %v", err)
		return &model.Stream{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      marshal_stream,
		TableName: aws.String(repo.table),
	}

	output, err := repo.client.PutItem(repo.ctx, input)

	if err != nil {
		log.Printf("Error while putting stream on CreateStream with dynamodb stream_table: %v", err)
		return &model.Stream{}, err
	}

	var stream_reponse model.Stream
	if err := attributevalue.UnmarshalMap(output.Attributes, &stream_reponse); err != nil {
		log.Printf("error unmarshaling CreateStream response: %v", err)
		return &model.Stream{}, err
	}

	return &stream_reponse, nil
}
