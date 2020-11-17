package action

import (
	"fmt"

	"github.com/deathwofl/cli-rpg-game/pkg"
)

// Train
func HitTheDummy(chr *pkg.Character) {

	chr.Strength++
	chr.Vitality++
	fmt.Printf("Max gained 1 Strength point.\n")
	fmt.Printf("Max gained 1 Vitality point.\n\n")
	UpdateStats(chr)
	PrintStats(chr)
}

func UpdateStats(chr *pkg.Character) {
	chr.HP = float32(chr.Vitality * uint16(10))
	chr.MP = float32(chr.Intelligence * 5)
}

func PrintStats(chr *pkg.Character) {
	fmt.Printf("HP:%v\n", chr.HP)
	fmt.Printf("MP:%v\n", chr.MP)
	fmt.Printf("Vitality:%v\n", chr.Vitality)
	fmt.Printf("Strength:%v\n", chr.Strength)
	fmt.Printf("Speed:%v\n", chr.Speed)
	fmt.Printf("Defense:%v\n", chr.Defense)
	fmt.Printf("Intelligence:%v\n", chr.Intelligence)
	fmt.Printf("Luck:%v\n", chr.Luck)
}
