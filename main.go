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
	// chatTable, err := services.CreateChatTable(client)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf(
	// 	"Table ID: %s \nTable Name: %s\n\n",
	// 	*chatTable.TableDescription.TableId,
	// 	*chatTable.TableDescription.TableName,
	// )

	// // add item to chart table
	// ChatID := uuid.New()
	// chatData := services.ChatDataType{
	// 	UserID:    "dbebf8e1-a375-4f9b-af6d-41f057e7b49b",
	// 	ChatID:    ChatID.String(),
	// 	Title:     "Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the blind texts.",
	// 	CreatedAt: time.Now().Unix(),
	// }

	// _, err = services.Create(client, chatData)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("New Chat Created.")

	chats, err := services.GetUserChats(client, "dbebf8e1-a375-4f9b-af6d-41f057e7b49b")

	if err != nil {
		fmt.Println(err)
	}

	// loop over the chats
	for _, chat := range *chats {
		fmt.Printf(
			"User ID: %s \nChat ID: %s \nChat Title: %s\n\n",
			chat.UserID,
			chat.ChatID,
			chat.Title,
		)
	}

	updatedFields, err := services.UpdateChat(client, services.ChatDataType{
		UserID: "dbebf8e1-a375-4f9b-af6d-41f057e7b49b",
		ChatID: "0955f724-e2fb-4787-8298-bcab5a159f0c",
		Title:  "The far away mountain loomed in the distance, its peak shrouded in mist and mystery.",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Updated %s \n\n", updatedFields.Attributes["chat_id"])

	// get single chat item
	chat, err := services.GetSingleChat(
		client,
		"dbebf8e1-a375-4f9b-af6d-41f057e7b49b",
		"0955f724-e2fb-4787-8298-bcab5a159f0c",
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(
		"\n\nSingle Chat \nUser ID: %s \nChat ID: %s \nChat Title: %s\n\n",
		chat.UserID,
		chat.ChatID,
		chat.Title,
	)

	// delete chat
	err = services.DeleteChat(
		client,
		"dbebf8e1-a375-4f9b-af6d-41f057e7b49b",
		"fb075ee7-f119-4a69-ac1f-0642c0441ca8",
	)

	if err != nil {
		fmt.Println(err)
	}
}
