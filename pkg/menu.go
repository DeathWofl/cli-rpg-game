package Pkg

type Menu struct {
	Choices  []string         // items on the to-do list
	Selected map[int]struct{} // which to-do items are selected
}
