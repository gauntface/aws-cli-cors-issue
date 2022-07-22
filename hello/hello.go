package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Hello world",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

type response struct {
	Message string `json:"message"`
}
