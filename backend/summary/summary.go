package summary

import (
	"github.com/Undefined01/fortuna/backend/importer"
)

type Summary struct {
	count   int
	tables  []importer.Table
	summary importer.TableMap
}

func New() *Summary {
	// 最终表格，第一个是汇总
	return &Summary{
		0,
		make([]importer.Table, 1, 20),
		importer.TableMap{
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

// 获取某一场考试某一班级的全部信息
// 参数：考试名称，班级（需带前置0）
// 返回：得分列表
func GetExamData(exam string, class string) []importer.Table {
	summary := New()
	urlPrefix := "http://192.168.206.6/20162018"
	termInfo := "year_in=2016&grade_name=%B8%DF%C8%FD&class_name1=2018-2019%C9%CF%D1%A7%C6%DA"
	eco := &importer.FromWebsite{
		ExamUrl:     []byte(urlPrefix + "/dk/top.asp?" + termInfo),
		ScoreUrl:    []byte(urlPrefix + "/dk/bottom_list_new.asp?" + termInfo),
		SubscoreUrl: []byte(urlPrefix + "/exam_a_p_g/bottom_list.asp?" + termInfo),
	}

/*
	urlPrefix := "http://192.168.1.88/2016abc"
	termInfo := "year_in=2016&grade_name=%B8%DF%C8%FD&class_name1=2018-2019%C9%CF%D1%A7%C6%DA"
	east := &importer.FromWebsite{
		ExamUrl:     []byte(urlPrefix + "/dk/top.asp?" + termInfo),
		ScoreUrl:    []byte(urlPrefix + "/dk/bottom_list_new.asp?" + termInfo),
		SubscoreUrl: []byte(urlPrefix + "/exam_a_p_g/bottom_list.asp?" + termInfo),
	}
*/

	sum := make(importer.ScoreMap)
	add := func (subject string) {
		score, subscore := eco.Get(exam, class, subject)
		summary.Add(subject, score, subscore)
		sum.Add(score)
	}
	// 添加科目
	add("语文")
	add("数学理")
	add("英语")
	add("物理")
	add("化学")
	add("生物")
	// 单独添加理科综合，防止重复计分
	score, subscore := eco.Get(exam, class, "理科综合")
	summary.Add("理科综合", score, subscore)

	summary.Add("总分", sum.Transform(), nil)

	return summary.Result()
}
