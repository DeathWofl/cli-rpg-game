package pkg

// Character describe fields
type Character struct {
	Name      string
	Money     uint64
	Stat      Stats
	Inventory map[int]Item
}

// CharacterService actions for character
type CharacterRepository interface {
	NewCharacter() *Character
	Cultivate(*Character) error
	Money(*Character) (*Character, error)
}

// NewCharacter create max character
func NewCharacter() Character {
	return Character{
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
