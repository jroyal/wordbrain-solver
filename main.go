package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jroyal/wordbrain-solver/grid"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var gridSize int
	var wordSize int
	var myGrid = grid.NewGrid()

	fmt.Print("Grid size: ")
	if _, err := fmt.Scan(&gridSize); err != nil {
		log.Print("  Grid Size input failed, due to ", err)
		return
	}

	for index := 0; index < gridSize; index++ {
		fmt.Printf("Grid row %d: ", index)
		gridLine, _ := reader.ReadString('\n')
		myGrid.AddRow(strings.Fields(gridLine))
	}

	fmt.Print("Size of word to solve for: ")

	if _, err := fmt.Scan(&wordSize); err != nil {
		log.Print("  Word size input failed, due to ", err)
		return
	}
	fmt.Print(myGrid.GetAllPossibleWords(wordSize))
}
