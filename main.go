package main

import "fmt"
import "github.com/jroyal/wordbrain-solver/grid"

func main() {
	fmt.Printf("Hello, world.\n")
	g := grid.NewGrid()
	g.AddRow([]string{"T", "E", "L", "H"})
	g.AddRow([]string{"R", "T", "C", "T"})
	g.AddRow([]string{"A", "I", "C", "I"})
	g.AddRow([]string{"H", "Y", "C", "W"})
	fmt.Println(g.GetAllPossibleWords(5))
}
