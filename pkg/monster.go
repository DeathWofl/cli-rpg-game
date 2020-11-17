package pkg

type Monster struct {
	Name string
	Stats
}

func NewMonster(name string, strenght uint16, vitality uint16, speed uint16, defense uint16) Monster {
	return Monster{
		Name: name,
		Stats: Stats{
			Vitality: vitality,
			Strength: strenght,
			Speed:    speed,
			Defense:  defense,
			HP:       float32(vitality * 10),
		},
	}
}
