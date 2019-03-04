package summary

import (
	"github.com/Undefined01/fortuna/backend/importer"
)

// 并行模块
type Parallel struct {
	Done  chan int
	Count int
}

func NewParallel() *Parallel {
	return &Parallel{
		Done:  make(chan int, 10),
		Count: 0,
	}
}
func (this *Parallel) Run(payload func()) {
	this.Count++
	hack := func() {
		payload()
		this.Done <- 1
	}
	go hack()
}
func (this *Parallel) Join() {
	for this.Count > 0 {
		<-this.Done
		this.Count--
	}
}

// 获取某一场考试某一班级全部科目的成绩
// 参数：考试名称，班级（需带前置0）
// 返回：得分列表
func GetExamData(exam string, class string) []importer.Table {
	summary := New()
	eco := importer.NewFromWebsite()

	subjectList := []string{"语文", "数学理", "英语", "物理", "化学", "生物", "理科综合"}
	type Result struct {
		Score    importer.ScoreList
		Subscore *importer.Subscore
	}
	result := make(map[int]Result, 10)
	sum := make(importer.RawScore)
	parallel := NewParallel()

	for i, _ := range subjectList {
		id := i
		parallel.Run(func() {
			score, subscore := eco.Get(exam, class, subjectList[id])
			result[id] = Result{score, subscore}
		})
	}
	parallel.Join()

	for i, _ := range subjectList {
		v, ok := result[i]
		if ok {
			if i != 6 {
				sum.Add(v.Score)
			}
			summary.Add(subjectList[i], v.Score, v.Subscore)
		}
	}
	summary.Add("总分", sum.Transform(), nil)

	return summary.Result()
}
