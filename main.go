package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type event struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

type response struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handleRequest(ctx context.Context, evt event) (string, error) {
	log.Printf("First Name: %s\n", evt.FirstName)
	log.Printf("Last Name: %s\n", evt.LastName)
	log.Printf("Age: %d\n", evt.Age)
	response := response{
		Name: fmt.Sprintf("%s %s", evt.FirstName, evt.LastName),
		Age:  evt.Age,
	}
	byt, _ := json.Marshal(response)
	resp := string(byt)
	log.Println(resp)
	return resp, nil
}

func main() {
	lambda.Start(handleRequest)
}
