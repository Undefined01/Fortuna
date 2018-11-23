package importer

// 数据导入模块：从学校的成绩查询网站获取成绩。
type FromWebsite struct {
	core *FromWebsiteCore
}

func NewFromWebsite() *FromWebsite {
	/*
		东城配置
		urlPrefix := "http://192.168.1.88/2016abc"
		termInfo := "year_in=2016&grade_name=%B8%DF%C8%FD&class_name1=2018-2019%C9%CF%D1%A7%C6%DA"
	*/
	urlPrefix := "http://192.168.206.6/20162018"
	termInfo := "year_in=2016&grade_name=%B8%DF%C8%FD&class_name1=2018-2019%C9%CF%D1%A7%C6%DA"
	core := &FromWebsiteCore{
		ExamUrl:     []byte(urlPrefix + "/dk/top.asp?" + termInfo),
		ScoreUrl:    []byte(urlPrefix + "/dk/bottom_list_new.asp?" + termInfo),
		SubscoreUrl: []byte(urlPrefix + "/exam_a_p_g/bottom_list.asp?" + termInfo),
	}
	return &FromWebsite{core}
}

func (this *FromWebsite) GetExamList() []string {
	return this.core.GetExamList()
}

// 获取总分和小题分成绩，并重整小题分数据。
func (this *FromWebsite) Get(exam string, class string, subject string) (ScoreList, *Subscore) {
	score := this.core.GetScore(exam, class, subject)
	subscore := this.core.GetSubscore(exam, class, subject)

	if score == nil || subscore == nil {
		return score, subscore.Transform(subject)
	}
	subscore.Cols = append(subscore.Cols, "主观得分", "客观得分", "总分", "级排名")
	for _, v := range score {
		p, ok := subscore.Data[v.Sid]
		if !ok {
			continue
		}
		subscore.Data[v.Sid] = append(p, v.Object, v.Subject, v.Total, v.Rank)
	}

	return score, subscore.Transform(subject)
}
