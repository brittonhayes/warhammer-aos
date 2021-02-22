package handlers

import (
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
)

var _ Handler = &handle{}

// Handler contains all routes for API request
// handling
type Handler interface {
	AddRoute(path string, handlers ...fiber.Handler) *handle
}

type handle struct {
	Application *fiber.App
}

// New creates a request handler
func New() *handle {
	return &handle{
		Application: fiber.New(),
	}
}

// AddRoute adds a fiber route to the request handler
func (h *handle) AddRoute(path string, handlers ...fiber.Handler) *handle {
	h.Application.Add("get", path, handlers...)
	return h
}

// Build constructs the fiber application into a
// lambda-compatible adapter
func (h *handle) Build() *fiberadapter.FiberLambda {
	return fiberadapter.New(h.Application)
}
