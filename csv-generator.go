package tocsv

import (
	"fmt"
	"io"
	"strings"
	"encoding/csv"
)

const (
	utf8_bom = "\xef\xbb\xbf"
)

type CSVGenerator interface {
	/// 在输出整个csv之前调用，在这里可以做一些输出准备工作
	BeforeOutputCSV()

	/// 获取输出目标
	GetWriter() io.Writer

	/// 获取csv的标题栏
	GetTitles() []string

	// 获取所有的输出行channel
	GetRows() (<-chan map[string]string)
}

func GenerateCSV(cg CSVGenerator) {
	cg.BeforeOutputCSV()

	rowsHandled := false
	rows := cg.GetRows()
	if rows == nil {
		return
	}
	defer func() {
		if rowsHandled {
			return
		}
		// rows必须读完，防止channel堵塞
		for _ = range rows {}
	}()

	writer := cg.GetWriter()
	if writer == nil {
		return
	}

	titles := cg.GetTitles()
	if len(titles) == 0 {
		return
	}

	fCsv := csv.NewWriter(writer)
	defer fCsv.Flush()
	outputTitles(fCsv, titles)
	rowsHandled = true

	row := make([]string, len(titles))
	for d := range rows {
		outputRow(fCsv, d, row, titles)
	}
}

func outputTitles(fCsv *csv.Writer, titles []string) {
	oldTitle0 := titles[0]
	hasBOM := strings.HasPrefix(oldTitle0, utf8_bom)
	if !hasBOM {
		titles[0] = fmt.Sprintf("%s%s", utf8_bom, oldTitle0)
	}

	fCsv.Write(titles)
	if !hasBOM {
		titles[0] = oldTitle0
	}
}

func outputRow(fCsv *csv.Writer, row map[string]string, outRow []string, titles []string) {
	for i, title := range titles {
		if col, ok := row[title]; ok {
			outRow[i] = col
		} else {
			outRow[i] = ""
		}
	}
	fCsv.Write(outRow)
}
