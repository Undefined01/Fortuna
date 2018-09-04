package server

import (
	"github.com/Undefined01/fortuna/backend/importer"
	"github.com/Undefined01/fortuna/backend/utils"
)

type DataGetter struct {
	// 配置
	class string
	// 缓存
	// 表格列表
	data []*utils.TableData
	// 表格总数
	count int
	// 总成绩表格的表头
	summaryCols []string
	// 总成绩表格数据的缓存，由于要统计同一个人的多个科目，使用学号作为键方便合并各科数据
	summaryTemp map[int][]interface{}
	sumTemp map[int]float32
}

func (this *DataGetter) Init() {
	// 最终表格，第一个是汇总
	this.data = make([]*utils.TableData, 1, 20)
	this.summaryCols = append(make([]string, 0, 20), "学号", "姓名")
	this.summaryTemp = make(map[int][]interface{})
	this.sumTemp = make(map[int]float32)
}

func (this *DataGetter) check(sid int, name string) {
	// 判断此人是否已有一行数据
	_, ok := this.summaryTemp[sid]
	if !ok {
		this.summaryTemp[sid] = append(make([]interface{}, 0, 20), sid, name)
	}
	// 补充未考数据
	for len(this.summaryTemp[sid]) < this.count*2 {
		this.summaryTemp[sid] = append(this.summaryTemp[sid], -1)
	}
}

func (this *DataGetter) Add(exam string, subject string, addToSum bool) []utils.ScoreOfSubject {
	fromWebsite := importer.NewFromWebsiteImporter()
	score := fromWebsite.GetScore(exam, this.class, subject)
	subscore := fromWebsite.GetSubscore(exam, this.class, subject)
	
	// 各个科目主客观得分缓存
	scoreTemp := make(map[int][3]float32)
	if score != nil {
		this.count++
		this.summaryCols = append(this.summaryCols, subject, subject+"排名")
		for _, person := range score {
			this.check(person.Sid, person.Name)
			this.summaryTemp[person.Sid] = append(this.summaryTemp[person.Sid], person.Total, person.Rank)
			if addToSum && person.Total != -1 {
				this.sumTemp[person.Sid] = this.sumTemp[person.Sid] + person.Total
			}
			scoreTemp[person.Sid] = [3]float32{person.Object, person.Subject, person.Total}
		}
		if subscore != nil {
			subscore.Cols = append(subscore.Cols, "主观得分", "客观得分", "总分")
			for i, _ := range subscore.Data {
				sid, ok := subscore.Data[i][0].(int)
				if !ok {
					continue
				}
				score, ok := scoreTemp[sid]
				if !ok {
					continue
				}
				subscore.Data[i] = append(subscore.Data[i], score[0], score[1], score[2])
			}
			this.data = append(this.data, subscore)
		}
	}
	return score
}

func (this *DataGetter) GetResult() []*utils.TableData {
	if this.count == 0 {
		return nil
	}
	var summary utils.TableData
	summary.Title = "总分"
	summary.Cols = append(this.summaryCols, "总分")
	summary.Data = make([][]interface{}, 0, len(this.summaryTemp))
	for k, v := range this.summaryTemp {
		summary.Data = append(summary.Data, append(v, this.sumTemp[k]))
	}
	this.data[0] = &summary
	return this.data
}

// 获取某一场考试某一班级的全部信息
// 参数：考试名称，班级（需带前置0）
// 返回：得分列表
func GetExamData(exam string, class string) []*utils.TableData {
	getter := &DataGetter{class: class}
	getter.Init()
	
	getter.Add(exam, "语文", true)
	getter.Add(exam, "数学", true)
	getter.Add(exam, "数学文", true)
	getter.Add(exam, "数学理", true)
	getter.Add(exam, "英语", true)
	
	getter.Add(exam, "物理", true)
	getter.Add(exam, "化学", true)
	getter.Add(exam, "生物", true)
	getter.Add(exam, "理科综合", false)
	
	getter.Add(exam, "政治", true)
	getter.Add(exam, "历史", true)
	getter.Add(exam, "地理", true)
	getter.Add(exam, "文科综合", false)

	
	return getter.GetResult()
}
