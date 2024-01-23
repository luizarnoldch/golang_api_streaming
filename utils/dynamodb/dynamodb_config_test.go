package dynamodb

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func TestGetDynamoDBAWSClient(t *testing.T) {
	ctx := context.TODO()
	client, err := GetDynamoDBAWSClient(ctx)
	assert.NoError(t, err)
	assert.IsType(t, &dynamodb.Client{}, client)
}

func TestGetDynamoDBAWSClientError(t *testing.T) {
	ctx := context.TODO()
	client, err := GetDynamoDBAWSClient(ctx)
	assert.NoError(t, err)
	assert.IsType(t, &dynamodb.Client{}, client)
}

func TestGetLocalDynamoDBClient(t *testing.T) {
	ctx := context.TODO()
	client, err := GetLocalDynamoDBClient(ctx)
	assert.NoError(t, err)
	assert.IsType(t, &dynamodb.Client{}, client)
	// Additional checks can be added to ensure that the client is configured for local use.
}

func TestGetLocalEndpoint(t *testing.T) {
	// Call the function
	endpoint, err := GetLocalEndpoint("dynamodb", "us-west-2")

	// Check for errors
	assert.NoError(t, err, "getLocalEndpoint should not return an error")

	// Check if the returned endpoint is as expected
	assert.Equal(t, "http://localhost:8000", endpoint.URL, "Endpoint URL should be 'http://localhost:8000'")
}