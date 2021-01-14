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
	fCsv := csv.NewWriter(cg.GetWriter())
	cg.BeforeOutputCSV()

	titles := cg.GetTitles()
	if len(titles) > 0 {
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

	row := make([]string, len(titles))
	for d := range cg.GetRows() {
		cg.BeforeOutputRow(d)
		for i, title := range titles {
			row[i] = cg.GetColValue(d, i, title)
		}
		fCsv.Write(row)
	}
	fCsv.Flush()
}
