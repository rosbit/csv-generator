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

func (t *csvTest) GetRows() (<-chan map[string]string) {
	rows := make(chan map[string]string)
	go func() {
		for i := 0; i < 10; i++ {
			row := make(map[string]string)
			for j := 0; j < 3; j++ {
				row[fmt.Sprintf("%c", 'a'+j)] = fmt.Sprintf("%c_%d_%d", 'a'+j, i+1, j+1)
			}
			rows <- row
		}
		close(rows)
	}()

	return rows
}
