package grid

import (
	"bufio"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Grid struct {
	grid    [][]string
	size    int
	results []string
	dict    map[string]bool
}

func NewGrid() *Grid {
	g := new(Grid)
	g.dict = loadDict()
	return g
}

func loadDict() map[string]bool {
	dict := map[string]bool{}
	pwd, _ := os.Getwd()
	pwd = pwd[:strings.LastIndex(pwd, "wordbrain-solver")+len("wordbrain-solver")]
	absPath, _ := filepath.Abs(path.Join(pwd, "data/dictionary.txt"))
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dict[strings.ToUpper(scanner.Text())] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dict
}

func (g *Grid) AddRow(r []string) {
	g.grid = append(g.grid, r)
	g.size = len(r)
}

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

func (g *Grid) GetAllPossibleWords(length int) []string {
	var current []string

	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			recursiveWalk(g, length, i, j, current, generateCleanWalkPath(g.size))
		}
	}
	return g.results
}

func (g *Grid) validMove(x int, y int, walkedPath [][]bool) bool {
	if x >= g.size || y >= g.size || x < 0 || y < 0 {
		return false
	}

	if walkedPath[y][x] {
		return false
	}

	return true
}

func existingResult(list []string, s string) bool {
	for _, elem := range list {
		if elem == s {
			return true
		}
	}
	return false
}

func recursiveWalk(g *Grid, length int, x int, y int, current []string, walkedPath [][]bool) {
	// fmt.Printf("x:%d y:%d current: %v walkedPath: %v\n", x, y, current, walkedPath)
	current = append(current, g.grid[y][x])
	walkedPath[y][x] = true
	if len(current) == length {
		currentString := strings.Join(current, "")
		if _, ok := g.dict[currentString]; ok && !existingResult(g.results, currentString) {
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
