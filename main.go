package main

import (
	"fmt"
	"go-dynamodb/initializers"
	"go-dynamodb/utils"
)

func main() {
	initializers.LoadEnvVariables()

	table, err := utils.DescribeTable("Messages")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(
		"Table ID: %s \nTable Name: %s\n\n",
		*table.Table.TableId, //as DescribeTable returns a pointer, we need to dereference the values
		*table.Table.TableName,
	)

}
