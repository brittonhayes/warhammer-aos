package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	warhammer_aos "warhammer-aos"
)

//
// // ArmiesHandler is our lambda handler invoked by the `lambda.Start` function call
// func ArmiesHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayV2HTTPResponse, error) {
// 	// return request.New(map[string]interface{}{
// 	// 	"example": req.Headers,
// 	// }).Reply()
// 	return events.APIGatewayV2HTTPResponse{
// 		StatusCode:      200,
// 		Body:            req.Body,
// 		IsBase64Encoded: false,
// 	}, nil
// }
//
// func main() {
// 	lambda.Start(ArmiesHandler)
// }

var fiberLambda *fiberadapter.FiberLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Fiber cold start")
	r := fiber.New()
	r.Get("/armies", func(ctx *fiber.Ctx) error {
		return ctx.JSON(warhammer_aos.GenJSON())
	})

	fiberLambda = fiberadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
