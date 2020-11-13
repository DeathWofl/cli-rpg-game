package main

import (
	"fmt"
	"os"

	"github.com/deathwofl/cli-rpg-game/pkg/menu"
)

func main() {
	if err := menu.RunGame(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
