package summary // 成绩汇总模块

import (
	"github.com/Undefined01/fortuna/backend/importer"
)

// 成绩汇总核心
type Summary struct {
	count   int
	tables  []importer.Table
	summary importer.RawTable
}

func New() *Summary {
	// 最终表格，第一个是汇总
	return &Summary{
		count:  0,
		tables: make([]importer.Table, 1, 20),
		summary: importer.RawTable{
			append(make([]string, 0, 20), "学号", "姓名"),
			make(map[int][]interface{}),
		},
	}
}

func (this *Summary) Add(subject string, score importer.ScoreList, subscore *importer.Subscore) {
	if score != nil {
		this.count++
		this.summary.Cols = append(this.summary.Cols, subject, subject+"排名")
		for _, person := range score {
			// 判断此人是否已有一行数据
			_, ok := this.summary.Data[person.Sid]
			if !ok {
				this.summary.Data[person.Sid] = append(make([]interface{}, 0, 20), person.Sid, person.Name)
			}
			// 补充缺失数据
			for len(this.summary.Data[person.Sid]) < this.count*2 {
				this.summary.Data[person.Sid] = append(this.summary.Data[person.Sid], -1)
			}
			// 添加本场考试数据
			this.summary.Data[person.Sid] = append(this.summary.Data[person.Sid], person.Total, person.Rank)
		}
	}
	if subscore != nil {
		this.tables = append(this.tables, importer.Table(*subscore))
	}
}

func (this *Summary) Result() []importer.Table {
	if this.count == 0 {
		return nil
	}

	this.tables[0] = this.summary.Transform("总分")
	return this.tables
}
