package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/brittonhayes/warhammer-aos/internal/handlers"
	"github.com/brittonhayes/warhammer-aos/pkg/armies"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
	fiberLambda = handlers.New().
		AddRoute("*/:name", armies.Find()).
		AddRoute("*/", armies.List()).
		Build()
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(ctx.Value("Root"), req)
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
