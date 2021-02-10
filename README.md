# csv生成器

 1. 通过接口实现csv输出与数据生成逻辑的分离
 1. 通过列名输出数据，免去列下表数数容易出错的问题
 1. 支持csv BOM头信息的输出，在各种系统上查看都不出现乱码
 1. csv生成文件还是通过网络输出，使用场景自主决定
 1. 通过提供adapter实现通用接口函数，减少实现代码

## 接口定义

```go
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
```

## 例子

```go
package main

import (
	"github.com/rosbit/csv-generator"
	"os"
	"io"
	"fmt"
)

func main() {
	cg := &csvTest{}
	tocsv.GenerateCSV(cg);
}

// ---- CSVGenerator implementation ----
type csvTest struct {
	tocsv.DummyCSVGeneratorAdapter
}

func (a *csvTest) BeforeOutputCSV() {
	fmt.Fprintf(os.Stderr, "BeforeOutputCSV() called\n")
}

func (a *csvTest) GetWriter() io.Writer {
	return os.Stdout
}

func (a *csvTest) GetTitles() []string {
	return []string{"a", "b", "c"}
}

func (a *csvTest) GetRows() (<-chan map[string]string) {
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
```

## 运行结果
```csv
a,b,c
11,12,13
21,22,23
31,32,33
41,42,43
51,52,53
61,62,63
71,72,73
81,82,83
91,92,93
101,102,103
```

