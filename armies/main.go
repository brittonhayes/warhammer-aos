package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"warhammer-aos/internal/handlers"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
	logrus.Info("Fiber cold start")
	r := fiber.New()
	r.Get("*", func(ctx *fiber.Ctx) error {
		logrus.Info("path=", ctx.Path())
		return ctx.JSON(handlers.Armies())
	})

	fiberLambda = fiberadapter.New(r)
}

func ArmiesHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(ArmiesHandler)
}
