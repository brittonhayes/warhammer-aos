package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/brittonhayes/warhammer-aos/internal/handlers"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
	fiberLambda = handlers.New().
		AddRoute("*", handlers.Armies()).
		Build()
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
