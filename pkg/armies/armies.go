// package armies contains the types, methods, and interfaces for interacting with Warhammer Age of Sigmar Armies
package armies

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	warhammer "github.com/brittonhayes/warhammer-aos"
	"github.com/brittonhayes/warhammer-aos/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/kennygrant/sanitize"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

var (
	_ Service = &Army{}
	_ Service = &Armies{}
)

var ErrNotFound = "resource not found"

// Service covers all available methods
// of the Armies package
type Service interface {
	reply() *handlers.Response
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
	Name          string          `json:"name"`
	Size          string          `json:"size"`
	Move          string          `json:"move"`
	Save          string          `json:"save"`
	Bravery       string          `json:"bravery"`
	Wounds        string          `json:"wounds"`
	MissileWeapon []MissileWeapon `json:"missile_weapon,omitempty"`
	MeleeWeapon   []MissileWeapon `json:"melee_weapon,omitempty"`
	Abilities     []Ability       `json:"abilities"`
	Keywords      []string        `json:"keywords"`
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
	var wg sync.WaitGroup

	// Iterate over all files in armies folder
	// and join them into one slice
	for _, file := range folder {
		wg.Add(1)
		file := file

		// Open a new goroutine for each
		go func() {
			b, err := warhammer.Files.ReadFile(warhammer.DataDir + "/" + file.Name())
			if err != nil {
				panic(errors.Wrap(err, "failed to read file"))
			}

			var a Army
			err = json.Unmarshal(b, &a)
			if err != nil {
				panic(errors.Wrap(err, "failed to parse JSON for unit"))
			}
			armies = append(armies, a)
			wg.Done()
		}()
	}

	wg.Wait()
	return armies
}

func (a Army) find(name string) (*Army, error) {
	name = strings.ReplaceAll(name, "/", "")
	name = sanitize.Path(name)
	b, err := warhammer.Files.ReadFile(warhammer.DataDir + "/" + name + ".json")
	if err != nil {
		err = errors.Wrapf(err, "failed to find %s", name)
		return nil, err
	}

	err = json.Unmarshal(b, &a)
	if err != nil {
		err = errors.Wrap(err, "failed to parse JSON for unit")
		panic(err)
	}

	return &a, nil
}

func (a *Army) unit(name string) string {
	b, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	return gjson.Get(string(b), fmt.Sprintf(`units.#(name=="%s")#`, name)).String()
}

// reply returns an army formatted as an
// API response
func (a *Army) reply() *handlers.Response {
	return &handlers.Response{
		Count: 1,
		Data:  a,
	}
}

// Reply returns a list of armies formatted
// as an API response
func (a *Armies) reply() *handlers.Response {
	return &handlers.Response{
		Count: len(*a),
		Data:  a,
	}
}

// List returns an http response all of the
// armies as a JSON
func List() func(ctx *fiber.Ctx) error {
	armies := readJSON()
	response := armies.reply()
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(response)
	}
}

// Find returns an http response all of the
// armies as a JSON
func Find() func(ctx *fiber.Ctx) error {

	var a Army
	return func(ctx *fiber.Ctx) error {
		name := ctx.Params("name")
		log.Println("Attempting to find ", name)
		army, err := a.find(name)
		if err != nil {
			return ctx.SendString(ErrNotFound)
		}

		return ctx.JSON(army.reply())
	}
}
