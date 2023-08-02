package main

import (
	"fmt"
	"go-dynamodb/constants"
	"go-dynamodb/infrastructure"
	"go-dynamodb/initializers"
	"go-dynamodb/services"
	"go-dynamodb/utils"
)

func main() {
	initializers.LoadEnvVariables()

	// load AWS config and dynamodb client
	config := infrastructure.NewAWSConfig()
	client := infrastructure.NewDynamoDBClient(config)

	table, err := utils.DescribeTable(client, constants.MessageTable)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(
		"Table ID: %s \nTable Name: %s\n\n",
		*table.Table.TableId, //as DescribeTable returns a pointer, we need to dereference the values
		*table.Table.TableName,
	)

	// create chats table
	chatTable, err := services.CreateChatTable(client)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(
		"Table ID: %s \nTable Name: %s\n\n",
		*chatTable.TableDescription.TableId,
		*chatTable.TableDescription.TableName,
	)

}
