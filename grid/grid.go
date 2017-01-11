package grid

import (
	"fmt"
	"strings"

	"github.com/jroyal/wordbrain-solver/dictionary"
)

// Grid is an object to hold the game state while being solved
type Grid struct {
	Grid    [][]string
	Size    int
	Words   []int
	results []string
}

// NewGrid will create a new Grid object, and build the dictionary used to check for correct words
func NewGrid() *Grid {
	g := new(Grid)
	dictionary.LoadDict()
	return g
}

// AddRow will add a row to the Grid object
func (g *Grid) AddRow(r []string) {
	g.Grid = append(g.Grid, r)
	g.Size = len(r)
}

// GetAllPossibleWords of a given length for the current grid
func (g *Grid) GetAllPossibleWords(length int) []string {
	fmt.Printf("Getting all words sized %d\n", length)
	var current []string

	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			recursiveWalk(g, length, i, j, current, generateCleanWalkPath(g.Size))
		}
	}
	return g.results
}

// generateCleanWalkPath creates a clean [][]bool array represention of the grid
func generateCleanWalkPath(size int) [][]bool {
	walkedPath := make([][]bool, size)
	for i := 0; i < size; i++ {
		walkedPath[i] = make([]bool, size)
		for j := 0; j < size; j++ {
			walkedPath[i][j] = false
		}
	}
	return walkedPath
}

// validMove checks to see if the requested position is a valid position to move to
func (g *Grid) validMove(x int, y int, walkedPath [][]bool) bool {
	if x >= g.Size || y >= g.Size || x < 0 || y < 0 {
		return false
	}

	if walkedPath[y][x] {
		return false
	}

	return true
}

// existingResult checks to see if the new solution already exists
func existingResult(list []string, s string) bool {
	for _, elem := range list {
		if elem == s {
			return true
		}
	}
	return false
}

// recursiveWalk takes a grid, a requested length and solves the grid for all words of that length
func recursiveWalk(g *Grid, length int, x int, y int, current []string, walkedPath [][]bool) {
	// fmt.Printf("x:%d y:%d current: %v walkedPath: %v\n", x, y, current, walkedPath)
	current = append(current, g.Grid[y][x])
	walkedPath[y][x] = true
	if len(current) == length {
		currentString := strings.Join(current, "")
		if dictionary.CheckWord(currentString) && !existingResult(g.results, currentString) {
			g.results = append(g.results, currentString)
		}
		return
	}

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if g.validMove(x+i, y+j, walkedPath) {
				recursiveWalk(g, length, x+i, y+j, current, walkedPath)
				walkedPath[y+j][x+i] = false
			}
		}
	}
	return
}
