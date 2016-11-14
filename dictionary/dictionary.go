package dictionary

import (
	"bufio"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var dict = map[string]bool{}

func LoadDict() {
	pwd, _ := os.Getwd()
	pwd = pwd[:strings.LastIndex(pwd, "wordbrain-solver")+len("wordbrain-solver")]
	absPath, _ := filepath.Abs(path.Join(pwd, "dictionary/dictionary.txt"))
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
}

func CheckWord(word string) bool {
	_, exists := dict[word]
	return exists
}
