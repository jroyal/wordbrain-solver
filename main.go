package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jroyal/wordbrain-solver/grid"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var gridSize int
	var wordSizes []int
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

	fmt.Print("Space separated lengths to solve for: ")

	wordSizeStr, _ := reader.ReadString('\n')
	for _, size := range strings.Fields(wordSizeStr) {
		i, _ := strconv.Atoi(size)
		wordSizes = append(wordSizes, i)
	}
	fmt.Print(myGrid.Solve(wordSizes))
}
