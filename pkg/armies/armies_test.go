package armies_test

import (
	"fmt"
	"testing"

	"github.com/brittonhayes/warhammer-aos/internal/handlers"
	"github.com/brittonhayes/warhammer-aos/pkg/armies"
)

// Create a new handler
// and add a function to the route
func ExampleArmies() {
	_ = handlers.New().AddRoute("*", armies.Handler()).Build()
}

func TestArmy_Unit(t *testing.T) {
	a := &armies.Army{
		Army: "example_name",
		Units: []armies.Unit{
			{Name: "first unit", Size: "100mmF"},
		},
	}

	result := a.Unit("first unit")
	fmt.Println(result)
}
