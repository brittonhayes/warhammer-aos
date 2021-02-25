package indexes

import (
	"github.com/brittonhayes/warhammer-aos/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

type Index string

const (
	RootIndex   Index = "/api"
	ArmiesIndex Index = "/api/armies"
	ArmiesParamIndex Index = "/api/armies/{name}"
)

// List lists out the the available API routes
func (i Index) List() []string {
	return []string{
		"/api",
		"/api/armies",
		"/api/armies/{name}",
	}
}

// Reply returns the indexes formatted as an
// API response
func (i *Index) Reply() *handlers.Response {
	results := i.List()
	return &handlers.Response{
		Count: len(results),
		Data:  results,
	}
}

// Handler returns an http response all of the
// indexes as a JSON
func Handler() func(ctx *fiber.Ctx) error {
	var index Index
	response := index.List()
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(response)
	}
}
