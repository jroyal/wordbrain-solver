package main

import "fmt"
import "github.com/jroyal/wordbrain-solver/grid"

func main() {
	fmt.Printf("Hello, world.\n")
	g := grid.NewGrid()
	g.AddRow([]string{"E", "A", "T", "I"})
	g.AddRow([]string{"G", "A", "V", "E"})
	g.AddRow([]string{"E", "E", "M", "A"})
	g.AddRow([]string{"L", "G", "N", "R"})
	fmt.Println(g.GetAllPossibleWords(8))
}
