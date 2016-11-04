package grid

import "testing"

func TestGetAllPossibleWordsGrid3Word4(t *testing.T) {
	g := NewGrid()
	g.AddRow([]string{"L", "S", "E"})
	g.AddRow([]string{"L", "I", "D"})
	g.AddRow([]string{"L", "O", "D"})
	expectedLength := 23
	result := g.GetAllPossibleWords(4)
	if len(result) != expectedLength {
		t.Logf("Expected: %d Result:%d", expectedLength, len(result))
		t.Fail()
	}
	found := false
	for _, answer := range result {
		if answer == "DOLL" {
			found = true
		}
	}
	if !found {
		t.Logf("DOLL should be in result slice %v", result)
		t.Fail()
	}
}

func TestGetAllPossibleWordsGrid3Word5(t *testing.T) {
	g := NewGrid()
	g.AddRow([]string{"L", "S", "E"})
	g.AddRow([]string{"L", "I", "D"})
	g.AddRow([]string{"L", "O", "D"})
	expectedLength := 8
	result := g.GetAllPossibleWords(5)
	if len(result) != expectedLength {
		t.Logf("Expected: %d Result:%d", expectedLength, len(result))
		t.Fail()
	}
	found := false
	for _, answer := range result {
		if answer == "SLIDE" {
			found = true
		}
	}
	if !found {
		t.Logf("DOLL should be in result slice %v", result)
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
