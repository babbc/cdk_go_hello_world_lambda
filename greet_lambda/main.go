package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Name string `json:"name"`
}

// Responds to JSON:
//
//	{
//		"name": "World"
//	}
func HandleRequest(ctx context.Context, event *Event) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	log.Printf("EVENT: %s", event.Name)

	message := fmt.Sprintf("Hello, %s!", event.Name)

	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
