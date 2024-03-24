package adapter

import (
	"context"
	"log"
	"main/src/users/domain/model"
	"main/src/users/domain/repository"
	appError "main/utils/error"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserDynamoDBRepository struct {
	client *dynamodb.Client
	ctx    context.Context
	table  string
}

func NewUserDynamoDBRepository(ctx context.Context, client *dynamodb.Client, table string) repository.UserRepository {
	return &UserDynamoDBRepository{
		ctx:    ctx,
		client: client,
		table:  table,
	}
}

func (repo *UserDynamoDBRepository) UpdateUserByID(userID string, user *model.User) (*model.User, *appError.Error) {
	updateBuilder := expression.Set(
		expression.Name("name"), expression.Value(user.Name)).Set(
		expression.Name("email"), expression.Value(user.Email))
	
	expr, err := expression.NewBuilder().WithUpdate(updateBuilder).Build()
	if err != nil {
		log.Printf("Error building expression for UpdateUserByID, %s", err.Error())
		return nil, appError.NewUnexpectedError(err.Error())
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(repo.table),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: userID},
		},
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ReturnValues:              types.ReturnValueUpdatedNew,
	}

	output, err := repo.client.UpdateItem(repo.ctx, input)
	if err != nil {
		log.Printf("Error updating stream UpdateUserByID, %s", err.Error())
		return nil, appError.NewUnexpectedError(err.Error())
	}
	var updatedUser model.User
	if err := attributevalue.UnmarshalMap(output.Attributes, &updatedUser); err != nil {
		log.Printf("Error unmarshaling UpdateUserByID, %s", err.Error())
		return nil, appError.NewUnexpectedError(err.Error())
	}
	log.Printf("UpdateUserByID done successfully")
	return &updatedUser, nil
}

func (repo *UserDynamoDBRepository) GetUserByID(userID string) (*model.User, *appError.Error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(repo.table),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: userID},
		},
	}
	result, err := repo.client.GetItem(repo.ctx, input)
	if err != nil {
		log.Printf("Error getting item from DynamoDB: %v", err)
		return nil, appError.NewUnexpectedError("Failed to get item from DynamoDB")
	}
	if result.Item == nil {
		log.Printf("No item found with on GetUserByID, %s", userID)
		return nil, appError.NewNotFoundError("User not found")
	}
	user := &model.User{}
	err = attributevalue.UnmarshalMap(result.Item, user)
	if err != nil {
		log.Printf("Error unmarshalling DynamoDB item to user: %v", err)
		return nil, appError.NewUnexpectedError("Failed to unmarshal DynamoDB item to user")
	}
	log.Printf("GetUserByID done successfully")
	return user, nil
}

func (repo *UserDynamoDBRepository) GetAllUser() ([]model.User, *appError.Error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(repo.table),
	}
	result, err := repo.client.Scan(repo.ctx, input)
	if err != nil {
		log.Printf("Error scanning DynamoDB table: %v", err)
		return nil, appError.NewUnexpectedError("Failed to scan DynamoDB table")
	}
	users := []model.User{}
	for _, item := range result.Items {
		user := model.User{}
		err := attributevalue.UnmarshalMap(item, &user)
		if err != nil {
			log.Printf("Error while unmarshaling item from GetAllUser, %s", err.Error())
			return nil, appError.NewUnexpectedError(err.Error())
		}
		users = append(users, user)
	}
	log.Printf("GetAllUser done successfully")
	return users, nil
}

func (repo *UserDynamoDBRepository) CreateUser(user *model.User) (*model.User, *appError.Error) {
	av, err := attributevalue.MarshalMap(user)
	if err != nil {
		log.Printf("Error marshalling user: %v", err)
		return nil, appError.NewUnexpectedError("Failed to marshal user")
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(repo.table),
	}
	_, err = repo.client.PutItem(repo.ctx, input)
	if err != nil {
		log.Printf("Error putting item to DynamoDB: %v", err)
		return nil, appError.NewUnexpectedError("Failed to put item into DynamoDB")
	}
	log.Printf("CreateUser done successfully")
	return user, nil
}

func (repo *UserDynamoDBRepository) DeleteUser(userID string) *appError.Error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(repo.table),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: userID},
		},
	}
	_, err := repo.client.DeleteItem(repo.ctx, input)
	if err != nil {
		log.Printf("Error deleting item from DynamoDB: %v", err)
		return appError.NewUnexpectedError("Failed to delete item from DynamoDB")
	}
	log.Printf("DeleteUser done successfully")
	return nil
}
