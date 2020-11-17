package action

import (
	"fmt"

	"github.com/deathwofl/cli-rpg-game/pkg"
)

// Cultivate get carrots
func Cultivate(chr *pkg.Character) {

	chr.Money += 15
	fmt.Println("Max gained 15 gold.")
	Money(chr)

}

// Money print character amount money
func Money(chr *pkg.Character) {
	fmt.Printf("Money: %v", chr.Money)
}
