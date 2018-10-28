// 成绩数据导入模块
package importer

type Importer interface {
	GetExamList() []string
	Get(exam string, class string, subject string) (ScoreList, *Subscore)
}
