package importer

// 数据导入模块：从学校的成绩查询网站获取成绩。
type FromWebsite struct {
	// 学校成绩查询网站的地址（包括年级信息）
	ExamUrl     []byte
	ScoreUrl    []byte
	SubscoreUrl []byte
}

func (this *FromWebsite) GetExamList() []string {
	core := &FromWebsiteCore{
		this.ExamUrl,
		this.ScoreUrl,
		this.SubscoreUrl,
	}
	return core.GetExamList()
}

func (this *FromWebsite) Get(exam string, class string, subject string) (ScoreList, *Subscore) {
	core := &FromWebsiteCore{
		this.ExamUrl,
		this.ScoreUrl,
		this.SubscoreUrl,
	}

	score := core.GetScore(exam, class, subject)
	subscore := core.GetSubscore(exam, class, subject)

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
