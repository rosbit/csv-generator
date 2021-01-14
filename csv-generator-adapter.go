package tocsv

import (
	"io"
	"os"
	"fmt"
)

type CSVGeneratorAdapter struct {
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

func (a *CSVGeneratorAdapter) GetRows() (<-chan interface{}) {
	rows := make(chan interface{})
	go func() {
		for i := 0; i < 10; i++ {
			row := make([]string, 3)
			for j := 0; j < 3; j++ {
				row[j] = fmt.Sprintf("%d%d", i+1, j+1)
			}
			rows <- row
		}

		close(rows)
	}()

	return rows
}

func (a *CSVGeneratorAdapter) BeforeOutputRow(row interface{}) {
	fmt.Fprintf(os.Stderr, "BeforeOutputRow() called, row: %v\n", row)
}

func (a *CSVGeneratorAdapter) GetColValue(row interface{}, idx int, title string) string {
	realRow, ok := row.([]string)
	if !ok {
		return ""
	}
	return realRow[idx]
}
