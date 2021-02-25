package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/brittonhayes/warhammer-aos/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

var fiberLambda *fiberadapter.FiberLambda

var robots = `
User-agent: *
Disallow: /
`

func init() {
	fiberLambda = handlers.New().
		AddRoute("*", func(ctx *fiber.Ctx) error {
			ctx.Set("Content-Type", "text/plain")
			ctx.Set("Content-Encoding", "UTF-8")
			return ctx.SendString(robots)
		}).
		Build()
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
