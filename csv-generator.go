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
	GetRows() (<-chan interface{})

	/// 在输出每行之前调用，在这里可以做一些判断条件收集
	BeforeOutputRow(row interface{})

	/// 获取某一列的值
	GetColValue(row interface{}, idx int, title string) string
}

func GenerateCSV(cg CSVGenerator) {
	cg.BeforeOutputCSV()

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

	rows := cg.GetRows()
	if rows == nil {
		return
	}

	row := make([]string, len(titles))
	for d := range rows {
		outputRow(cg, fCsv, d, row, titles)
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

func outputRow(cg CSVGenerator, fCsv *csv.Writer, row interface{}, outRow []string, titles []string) {
		cg.BeforeOutputRow(row)
		for i, title := range titles {
			outRow[i] = cg.GetColValue(row, i, title)
		}
		fCsv.Write(outRow)
}
