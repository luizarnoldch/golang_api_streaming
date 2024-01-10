package configuration

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var (
	STREAM_TABLE = os.Getenv("STREAM_TABLE")
)

func GetDynamoDBStreamTable() string {
	if STREAM_TABLE == "" {
		log.Printf("Local DynamoDB Database")
		return "Test_Stream_Table"
	}
	log.Printf("AWS DynamoDB Database")
	return STREAM_TABLE
}

func CreateLocalDynamoDBStreamTable(client *dynamodb.Client, ctx context.Context, tableName string) error {
	_, err := client.CreateTable(ctx, &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName:   aws.String(tableName),
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		log.Fatalf("Failed to create Local DynamoDB AWS SDK v2 config: %s", err)
		return err
	}
	log.Printf("Table %s created successfully", tableName)
	return nil
}

func DeleteLocalDynamoDBStreamTable(client *dynamodb.Client, ctx context.Context, tableName string) error {
	_, err := client.DeleteTable(ctx, &dynamodb.DeleteTableInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		log.Fatalf("Failed to delete Local DynamoDB AWS SDK v2 config: %s", err)
		return err
	}

	log.Printf("Table %s deleted successfully", tableName)
	return nil
}
