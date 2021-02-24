// package armies contains the types, methods, and interfaces for interacting with Warhammer Age of Sigmar Armies
package armies

import (
	"encoding/json"
	"fmt"
	"log"

	warhammer "github.com/brittonhayes/warhammer-aos"
	"github.com/brittonhayes/warhammer-aos/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

var (
	_ Service = &Army{}
	_ Service = &Armies{}
)

// Service covers all available methods
// of the Armies package
type Service interface {
	Reply() *handlers.Response
}

// Armies is a helper type
// to allow receiver methods for
// a slice of Army
type Armies []Army

// Army is the parent structure that all nested
// units belong to
type Army struct {
	Army  string `json:"army"`
	Units []Unit `json:"units"`
}

// Unit is an individual unit in Warhammer
type Unit struct {
	Name string `json:"name"`
	Size string `json:"size"`
	// TODO re-implement these fields after json inconsistencies are fixed

	// Move          string          `json:"move"`
	// Save          string          `json:"save"`
	// Bravery       string          `json:"bravery"`
	// Wounds        string          `json:"wounds"`
	// MissileWeapon []MissileWeapon `json:"missile_weapon,omitempty"`
	// MeleeWeapon   []MissileWeapon `json:"melee_weapon,omitempty"`
	// Abilities     []Ability       `json:"abilities"`
	// Keywords      []string        `json:"keywords"`
}

// MissileWeapon is a weapon that is used
// for ranged attacks
type MissileWeapon struct {
	Name    string `json:"name"`
	Range   string `json:"range"`
	Attacks string `json:"attacks"`
	ToHit   string `json:"to_hit"`
	ToWound string `json:"to_wound"`
	Rend    string `json:"rend"`
	Damage  string `json:"damage"`
}

// MeleeWeapon is a weapon that is used
// for close quarters combat
type MeleeWeapon struct {
	Name    string `json:"name"`
	Range   string `json:"range"`
	Attacks string `json:"attacks"`
	ToHit   string `json:"to_hit"`
	ToWound string `json:"to_wound"`
	Rend    string `json:"rend"`
	Damage  string `json:"damage"`
}

// Ability is a skill possessed by the
// parent unit
type Ability struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func readJSON() Armies {
	folder, err := warhammer.Files.ReadDir(warhammer.DataDir)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to read json directory"))
	}
	var armies []Army
	for _, file := range folder {
		b, err := warhammer.Files.ReadFile(warhammer.DataDir + "/" + file.Name())
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to read file"))
			return nil
		}

		var a Army
		err = json.Unmarshal(b, &a)
		if err != nil {
			// log.Fatal(errors.Wrap(err, "failed to parse JSON for unit"))
			// return nil
			panic(err)
		}
		armies = append(armies, a)
	}

	return armies
}

func (a *Army) Unit(name string) string {
	b, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	return gjson.Get(string(b), fmt.Sprintf(`units.#(name=="%s")#`, name)).String()
}

// Reply returns an army formatted as an
// API response
func (a *Army) Reply() *handlers.Response {
	return &handlers.Response{
		Count: 1,
		Data:  a.Units,
	}
}

// Reply returns a list of armies formatted
// as an API response
func (a *Armies) Reply() *handlers.Response {
	return &handlers.Response{
		Count: len(*a),
		Data:  a,
	}
}

// Handler returns an http response all of the
// armies as a JSON
func Handler() func(ctx *fiber.Ctx) error {
	armies := readJSON()
	response := armies.Reply()
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(response)
	}
}
