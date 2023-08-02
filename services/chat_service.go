package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func CreateChatTable(client *dynamodb.Client) (*dynamodb.CreateTableOutput, error) {
	// create chats table
	table, err := client.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName: aws.String("Chats"),
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("user_id"),
			AttributeType: types.ScalarAttributeTypeS,
		}, {
			AttributeName: aws.String("chat_id"),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		KeySchema: []types.KeySchemaElement{{
			AttributeName: aws.String("user_id"),
			KeyType:       types.KeyTypeHash,
		}, {
			AttributeName: aws.String("chat_id"),
			KeyType:       types.KeyTypeRange,
		}},
		BillingMode: types.BillingModePayPerRequest,
	})

	return table, err
}