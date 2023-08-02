package utils

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func DescribeTable(client *dynamodb.Client, tableName string) (*dynamodb.DescribeTableOutput, error) {

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
