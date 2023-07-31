package main

import (
	"context"
	"fmt"
	"go-dynamodb/infrastructure"
	"go-dynamodb/initializers"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	initializers.LoadEnvVariables()

	// load AWS config and dynamodb client
	config := infrastructure.NewAWSConfig()
	client := infrastructure.NewDynamoDBClient(config)

	// get table information
	table, err := client.DescribeTable(
		context.TODO(),
		&dynamodb.DescribeTableInput{
			TableName: aws.String("Messages"),
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(
		"Table ID: %s \nTable Name: %s\n\n",
		*table.Table.TableId, //as DescribeTable returns a pointer, we need to dereference the values
		*table.Table.TableName,
	)

}
