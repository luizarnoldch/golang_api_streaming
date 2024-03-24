package main

import (
	"context"
	"fmt"
	"log"
	stream_configuration "main/src/streams/infrastructure/configuration"
    user_configuration "main/src/users/infrastructure/configuration"
	dynamodbUtils "main/utils/dynamodb"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)


func CreateInitLocalDynamoDBStreamTable(ctx context.Context, client *dynamodb.Client, tableName string) error {
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
        return fmt.Errorf("failed to create table: %w", err)
    }

    // Wait until the table is created and active
    describeInput := &dynamodb.DescribeTableInput{
        TableName: aws.String(tableName),
    }

    err = dynamodb.NewTableExistsWaiter(client).Wait(ctx, describeInput, 10*time.Second)
    if err != nil {
        return fmt.Errorf("failed to wait for table to become active: %w", err)
    }

    log.Printf("Table %s created successfully", tableName)
    return nil
}


func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
    defer cancel()

    dynamoClient, err := dynamodbUtils.GetLocalDynamoDBClient(ctx)
    if err != nil {
        log.Fatalf("Failed to create DynamoDB client: %v", err)
    }

    streamTableName := "StreamTable"
    streamTableExists, err := stream_configuration.DescribeStreamTable(ctx, dynamoClient, streamTableName)
    if err != nil {
        log.Fatalf("Error describing table: %v", err)
    }

    if !streamTableExists {
        if err := CreateInitLocalDynamoDBStreamTable(ctx, dynamoClient, streamTableName); err != nil {
            log.Fatalf("Failed to create table: %v", err)
        }
    } else {
        log.Printf("Table %s already exists", streamTableName)
    }

    userTableName := "UserTable"
    userTableExists, err := user_configuration.DescribeUserTable(ctx, dynamoClient, userTableName)
    if err != nil {
        log.Fatalf("Error describing table: %v", err)
    }

    if !userTableExists {
        if err := CreateInitLocalDynamoDBStreamTable(ctx, dynamoClient, userTableName); err != nil {
            log.Fatalf("Failed to create table: %v", err)
        }
    } else {
        log.Printf("Table %s already exists", userTableName)
    }
}
