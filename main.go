package main

import (
	"fmt"
	"go-dynamodb/initializers"
	"os"
)

func main() {
	initializers.LoadEnvVariables()
	fmt.Println(os.Getenv("AWS_REGION"))
}
