package configuration_test

import (
	"context"
	"main/src/streams/infrastructure/configuration"
    dynamodbUtils "main/utils/dynamodb"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

func TestGetDynamoDBStreamTableWithEnvSet(t *testing.T) {
    // Establecer la variable de entorno para la prueba
    testValue := "MyStreamTable"
    os.Setenv("STREAM_TABLE", testValue)
    defer os.Unsetenv("STREAM_TABLE") // Limpiar después de la prueba

    // Llamar a la función y verificar el resultado
    result := configuration.GetDynamoDBStreamTable()
    assert.Equal(t, testValue, result, "El valor obtenido debe ser igual al valor de la variable de entorno")
}

func TestGetDynamoDBStreamTableWithoutEnvSet(t *testing.T) {
    // Asegurarse de que la variable de entorno no esté establecida
    os.Unsetenv("STREAM_TABLE")

    // Llamar a la función y verificar el resultado
    result := configuration.GetDynamoDBStreamTable()
    assert.Equal(t, "Test_Stream_Table", result, "Si STREAM_TABLE no está establecido, se debe devolver 'Test_Stream_Table'")
}

func TestCreateLocalDynamoDBStreamTable(t *testing.T) {
    ctx := context.TODO()
    client, err := dynamodbUtils.GetLocalDynamoDBClient(ctx)
    assert.NoError(t, err)

    tableName := configuration.GetDynamoDBStreamTable()

    _ = configuration.DeleteLocalDynamoDBStreamTable(client, ctx, tableName)

    err = configuration.CreateLocalDynamoDBStreamTable(client, ctx, tableName)
    assert.NoError(t, err)

    describeInput := &dynamodb.DescribeTableInput{
        TableName: aws.String(tableName),
    }
    _, err = client.DescribeTable(ctx, describeInput)
    assert.NoError(t, err, "La tabla debe existir en DynamoDB local")

    err = configuration.DeleteLocalDynamoDBStreamTable(client, ctx, tableName)
	assert.NoError(t, err)
}

func TestDeleteTableNotExists(t *testing.T) {
    ctx := context.TODO()
    client, err := dynamodbUtils.GetLocalDynamoDBClient(ctx)
    assert.NoError(t, err)

    tableName := "NonExistentTable"

    err = configuration.DeleteLocalDynamoDBStreamTable(client, ctx, tableName)

    assert.Error(t, err, "Expected an error when trying to delete a non-existent table")

    assert.Contains(t, err.Error(), "does not exist, no need to delete")
}
