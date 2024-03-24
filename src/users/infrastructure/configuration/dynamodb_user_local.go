package configuration

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetDynamoDBUserTable() string {
	userTable := os.Getenv("USER_TABLE")
	if userTable == "" {
		log.Printf("Local DynamoDB Database")
		return "Test_User_Table"
	}
	log.Printf("AWS DynamoDB Database: %s", userTable)
	return userTable
}

func CreateLocalDynamoDBUserTable(ctx context.Context, client *dynamodb.Client, tableName string) error {
	client.CreateTable(ctx, &dynamodb.CreateTableInput{
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
	log.Printf("Table %s created successfully", tableName)
	return nil
}

func DescribeUserTable(ctx context.Context, client *dynamodb.Client, tableName string) (bool, error) {
	_, err := client.DescribeTable(ctx, &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})

	if err != nil {
		var notFoundErr *types.ResourceNotFoundException
		if errors.As(err, &notFoundErr) {
			log.Printf("Table %s does not exist.", tableName)
			return false, nil
		}
		log.Printf("Unexpected error occurred while describing table %s: %s", tableName, err)
		return false, err
	}

	log.Printf("Table %s exists.", tableName)
	return true, nil
}

func DeleteLocalDynamoDBUserTable(ctx context.Context, client *dynamodb.Client, tableName string) error {
    _, err := client.DescribeTable(ctx, &dynamodb.DescribeTableInput{
        TableName: aws.String(tableName),
    })

    if err != nil {
        return fmt.Errorf("table %s does not exist, no need to delete", tableName)
    }
    client.DeleteTable(ctx, &dynamodb.DeleteTableInput{
        TableName: aws.String(tableName),
    })

    log.Printf("Table %s deleted successfully", tableName)
    return nil
}
