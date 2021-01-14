package tocsv

import (
	"testing"
	"fmt"
	"log"
)

func testCSV(cg CSVGenerator, prompt string) {
	GenerateCSV(cg)
	log.Printf("testing %s done!\n", prompt)
}

func TestCSVAdapter(t *testing.T) {
	cg := &CSVGeneratorAdapter{}
	testCSV(cg, "TestCSVAdapter")
}

func TestCSV(t *testing.T) {
	cg := &csvTest{}
	testCSV(cg, "TestCSV")
}

type csvTest struct {
	CSVGeneratorAdapter
}

func (t *csvTest) BeforeOutputCSV() {
}

func (t *csvTest) BeforeOutputRow(row interface{}) {
}

func (t *csvTest) GetColValue(row interface{}, idx int, title string) string {
	realRow, ok := row.([]string)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s_%d_%s", title, idx, realRow[idx])
}
