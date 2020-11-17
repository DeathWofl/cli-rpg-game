package action

import (
	"fmt"

	"github.com/deathwofl/cli-rpg-game/pkg"
)

func KillGoblins(chr *pkg.Character) {

	goblin := pkg.NewMonster("Goblin", 4, 3, 4, 2)
	Battle(chr, &goblin)
	PrintStats(chr)
}

func Battle(chr *pkg.Character, mon *pkg.Monster) {
	for mon.HP > 0 {
		Turn(chr, mon)
		if mon.HP <= 0 {
			fmt.Printf("%s died.\n", mon.Name)
			fmt.Printf("Max gained rest of %s\n", mon.Name)
			break
		} else if chr.HP <= 0 {
			fmt.Println("Max died.")
			break
		}
	}
}

func Turn(chr *pkg.Character, mon *pkg.Monster) {
	mon.HP = mon.HP - CalculateDamage(int(chr.Strength), int(mon.Defense))
	fmt.Printf("Monster HP %v\n", mon.HP)
	chr.HP = chr.HP - CalculateDamage(int(mon.Strength), int(chr.Defense))
	fmt.Printf("Character HP %v\n", chr.HP)
}

func CalculateDamage(strength, defense int) float32 {

	if (strength - defense) < 0 {
		return 0
	}
	return float32(strength - defense)
}
