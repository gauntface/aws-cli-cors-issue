package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gauntface/focus-api/utils/jsonresponse"
)

func handler() (events.APIGatewayProxyResponse, error) {
	return jsonresponse.Data(response{
		Message: "hello world",
	}), nil
}

func main() {
	lambda.Start(handler)
}

type response struct {
	Message string `json:"message"`
}
