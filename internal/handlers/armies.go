package handlers

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	warhammer "warhammer-aos"
)

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
	Move          interface{}     `json:"move"`
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

func readArmies() []Army {
	folder, err := warhammer.Files.ReadDir("data/json")
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to read json directory"))
	}
	var armies []Army
	for _, file := range folder {
		b, err := warhammer.Files.ReadFile("data/json/" + file.Name())
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

// Armies returns an http response all of the
// armies as a JSON
func Armies() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(readArmies())
	}
}
