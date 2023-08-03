package services

import (
	"context"
	"fmt"
	"go-dynamodb/constants"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ChatDataType struct {
	ChatID    string    `json:"chat_id" dynamodbav:"chat_id"`
	UserID    string    `json:"user_id" dynamodbav:"user_id"`
	Title     string    `json:"title" dynamodbav:"title"`
	CreatedAt time.Time `json:"created_at" dynamodbav:"created_at"`
}

func CreateChatTable(client *dynamodb.Client) (*dynamodb.CreateTableOutput, error) {
	// create chats table
	table, err := client.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName: aws.String(constants.ChatTable),
		// primary key attributes are required
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("user_id"),
			AttributeType: types.ScalarAttributeTypeS,
		}, {
			AttributeName: aws.String("chat_id"),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		// add primary key details
		KeySchema: []types.KeySchemaElement{{
			AttributeName: aws.String("user_id"),
			KeyType:       types.KeyTypeHash,
		}, {
			AttributeName: aws.String("chat_id"),
			KeyType:       types.KeyTypeRange,
		}},
		// set billing mode
		BillingMode: types.BillingModePayPerRequest,
	})

	return table, err
}

// Create creates a new entry in dynamodb Chats table
func Create(
	client *dynamodb.Client,
	chatData ChatDataType,
) (*dynamodb.PutItemOutput, error) {
	// convert data to dynamodb AttributeValue
	av, err := attributevalue.MarshalMap(chatData)
	if err != nil {
		fmt.Printf("Got error marshalling data: %s\n", err)
		return nil, err
	}
	// save chat to db
	output, err := client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(constants.ChatTable), Item: av,
	})
	if err != nil {
		fmt.Printf("Couldn't add item to table.: %v\n", err)
	}
	return output, err
}
