package utils

import (
	"context"
	"go-dynamodb/infrastructure"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func DescribeTable(tableName string) (*dynamodb.DescribeTableOutput, error) {
	// load AWS config and dynamodb client
	config := infrastructure.NewAWSConfig()
	client := infrastructure.NewDynamoDBClient(config)

	// get table information
	table, err := client.DescribeTable(
		context.TODO(),
		&dynamodb.DescribeTableInput{
			TableName: aws.String(tableName),
		},
	)

	// return table and error
	return table, err
}
