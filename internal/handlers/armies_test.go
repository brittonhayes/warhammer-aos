package handlers_test

import (
	"warhammer-aos/internal/handlers"
)

// Create a new handler
// and add a function to the route
func ExampleArmies() {
	_ = handlers.New().
		AddRoute("*", handlers.Armies()).
		Build()
}
