package menu

import (
	"fmt"

	"github.com/muesli/termenv"
)

// RunGame initizalize the game
func RunGame() error {

	out := termenv.String("Welcome to Grimhunt, a cli rpg game where you will live an adventure from your terminal!!\n your name is max, a young man from a very remote town on the edge of the continent. He goes on an adventure around the world in search of revenge for his mother who was killed by demons when he was 5 years old, training tirelessly, until he is old enough to explore the continent.")
	p := termenv.ColorProfile()

	// setting text color
	out = out.Foreground(p.Color("#9ba4b4"))

	fmt.Print("\n")
	fmt.Println(out)

	// ShowMenu()

	return nil
}

// func ShowMenu() {
// 	text := termenv.String("\nYour adventure begins today, what do you decide to do to start?")
// 	p := termenv.ColorProfile()

// 	text = text.Foreground(p.Color("#9ba4b4"))
// 	fmt.Println(text)

// 	var initialModel = models.Menu{
// 		Choices:  []string{"Kill goblins", "Train", "Work in the field."},
// 		Selected: make(map[int]struct{}),
// 	}

// }
