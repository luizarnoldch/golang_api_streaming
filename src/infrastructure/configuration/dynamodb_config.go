package configuration

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	STREAM_TABLE = os.Getenv("STREAM_TABLE")
)

func GetDynamoDBAWSClient(ctx context.Context) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal("Failed to load DynamoDB AWS SDK v2 config")
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
}

func GetDynamoDBStreamTable() string {
	if STREAM_TABLE == "" {
		log.Printf("Local DynamoDB Database")
		return "Test_Stream_Table"
	}
	return STREAM_TABLE
}