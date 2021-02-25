package armies_test

import (
	"github.com/brittonhayes/warhammer-aos/internal/handlers"
	"github.com/brittonhayes/warhammer-aos/pkg/armies"
)

// Create a new handler
// and add a function to the route
func ExampleArmies() {
	_ = handlers.New().AddRoute("*", armies.List()).Build()
}
