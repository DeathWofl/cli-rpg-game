package Pkg

// Character describe fields
type Character struct {
	Name  string
	Money uint64
	Stat  Stats
}

// CharacterService actions for character
type CharacterService interface {
	NewCharacter() *Character
	Cultivate() (Character, error)
}

// NewCharacter create max character
func NewCharacter() *Character {
	return &Character{
		Name:  "Max",
		Money: 0,
		Stat: Stats{
			HP:           100,
			MP:           60,
			Vitality:     10,
			Strength:     5,
			Speed:        6,
			Defense:      10,
			Intelligence: 12,
			Luck:         5,
		},
	}
}
