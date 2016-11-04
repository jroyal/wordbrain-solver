package grid

import (
	"reflect"
	"testing"
)

func TestGetAllPossibleWordsGrid3Word4(t *testing.T) {
	g := NewGrid()
	g.AddRow([]string{"L", "S", "E"})
	g.AddRow([]string{"L", "I", "D"})
	g.AddRow([]string{"L", "O", "D"})
	expected := []string{"LIDS", "LIED", "LIES", "LOIS", "LODE", "LOLL", "SIDE", "SILL", "SILO", "SLID",
		"IDES", "IDOL", "ILLS", "ODDS", "ODIS", "OISE", "OILS", "ODES", "DILL", "DIDO", "DIES", "DOLL", "DIED"}
	result := g.GetAllPossibleWords(4)
	if !reflect.DeepEqual(expected, result) {
		t.Fail()
	}
}

func TestGetAllPossibleWordsGrid3Word5(t *testing.T) {
	g := NewGrid()
	g.AddRow([]string{"L", "S", "E"})
	g.AddRow([]string{"L", "I", "D"})
	g.AddRow([]string{"L", "O", "D"})
	expected := []string{"LODES", "LOLLS", "SLIDE", "IDOLS", "OLLIE", "DILLS", "DOLLS", "DIODE"}
	result := g.GetAllPossibleWords(5)
	if !reflect.DeepEqual(expected, result) {
		t.Fail()
	}
}

func BenchmarkGetAllPossibleWords(b *testing.B) {
	g := NewGrid()
	g.AddRow([]string{"L", "S", "E"})
	g.AddRow([]string{"L", "I", "D"})
	g.AddRow([]string{"L", "O", "D"})
	for n := 0; n < b.N; n++ {
		g.GetAllPossibleWords(4)
	}
}
