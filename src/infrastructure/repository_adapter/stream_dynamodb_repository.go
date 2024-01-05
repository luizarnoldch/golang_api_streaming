package repositoryadapter

import (
	"context"
	"main/src/domain/model"

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
	return []model.Stream{}, nil
}
