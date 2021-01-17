package tocsv

import (
	"io"
	"os"
	"fmt"
)

// ------ DummyCSVGeneratorAdapter -----
type DummyCSVGeneratorAdapter struct {}
func (a *DummyCSVGeneratorAdapter) BeforeOutputCSV() {}
func (a *DummyCSVGeneratorAdapter) GetWriter() io.Writer { return nil; }
func (a *DummyCSVGeneratorAdapter) GetTitles() []string { return nil; }
func (a *DummyCSVGeneratorAdapter) GetRows() (<-chan map[string]string) { return nil; }

// ------ CSVGeneratorAdapter ------
type CSVGeneratorAdapter struct {
	DummyCSVGeneratorAdapter
}

func (a *CSVGeneratorAdapter) BeforeOutputCSV() {
	fmt.Fprintf(os.Stderr, "BeforeOutputCSV() called\n")
}

func (a *CSVGeneratorAdapter) GetWriter() io.Writer {
	return os.Stdout
}

func (a *CSVGeneratorAdapter) GetTitles() []string {
	return []string{"a", "b", "c"}
}

func (a *CSVGeneratorAdapter) GetRows() (<-chan map[string]string) {
	rows := make(chan map[string]string)
	go func() {
		for i := 0; i < 10; i++ {
			row := make(map[string]string)
			for j := 0; j < 3; j++ {
				row[fmt.Sprintf("%c", 'a'+j)] = fmt.Sprintf("%d%d", i+1, j+1)
			}
			rows <- row
		}

		close(rows)
	}()

	return rows
}

