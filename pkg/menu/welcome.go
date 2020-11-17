package menu

import (
	"fmt"
	"os"

	"github.com/deathwofl/cli-rpg-game/pkg"
	"github.com/deathwofl/cli-rpg-game/pkg/action"
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

	fmt.Println("Your adventure begins today,")

	// options of menu
	mm := pkg.Menu{
		Choices: []string{"Kill goblins", "Train", "Work in the field.", "Exit"},
	}

	prota := pkg.NewCharacter()

	ShowMenu(&mm, &prota)

	return nil
}

// ShowMenu
func ShowMenu(mm *pkg.Menu, chr *pkg.Character) {

	text := termenv.String("\nwhat do you decide to do to start ?")
	p := termenv.ColorProfile()

	text = text.Foreground(p.Color("#9ba4b4"))
	fmt.Println(text)

	for index, model := range mm.Choices {
		fmt.Printf("%v. %s\n", index+1, model)
	}

	var option string
	fmt.Scanln(&option)

	// clean console or terminal
	// fmt.Println("\033[2J")
	// cmd := exec.Command("cmd", "/c", "cls")
	// cmd.Stdout = os.Stdout
	// cmd.Run()

	switch option {
	case "1":
		action.KillGoblins(chr)
		ShowMenu(mm, chr)
	case "2":
		action.HitTheDummy(chr)
		ShowMenu(mm, chr)
	case "3":
		action.Cultivate(chr)
		ShowMenu(mm, chr)
	case "4":
		os.Exit(1)
	default:
		fmt.Println("Select correct option.")
	}

}
