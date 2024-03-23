package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	STREAM_TABLE = os.Getenv("STREAM_TABLE")
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{URL: "http://localhost:8000"}, nil
				},
			),
		),
	)
	if err != nil {
		log.Println("error on load default confing")
		return events.APIGatewayProxyResponse{
			Body: "error on load default confing",
			StatusCode: 500,
		}, err
	}

	dynamoDBClient := dynamodb.NewFromConfig(cfg)
	
	input := &dynamodb.ListTablesInput{}
	result, err := dynamoDBClient.ListTables(ctx, input)
	if err != nil {
		log.Println("error on list tables")
		return events.APIGatewayProxyResponse{
			Body: "error on list tables",
			StatusCode: 500,
		}, err
	}
	
	message := "Tablas de DynamoDB:\n"
	for _, table_read := range result.TableNames {
		message += fmt.Sprintf("- %s\n", table_read)
	}

	return events.APIGatewayProxyResponse{
		Body: message,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
