package main

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Event events.KafkaEvent

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event Event) error {
	if len(event.Records) == 0 {
		return errors.New("no MSK message passed to function")
	}

	for _, records := range event.Records {
		for _, record := range records {
			log.Printf("Test TOPIC: %s", record.Topic)
		}
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
