// 数据导入模块
package importer

import "github.com/Undefined01/fortuna/backend/utils"

type Importer interface {
	GetExamList() []string
	GetScore(exam string, class string, subject string) []utils.ScoreOfSubject
	GetSubscore(exam string, class string, subject string) *utils.TableData
}

func NewFromWebsiteImporter() Importer {
	/*
	生态高二下
		termInfo := "year_in=2016&grade_name=%B8%DF%C8%FD&class_name1=2018-2019%C9%CF%D1%A7%C6%DA"
		// 2017-2018%CF%C2%D1%A7%C6%DA
		return &FromWebsite{
			ExamUrl:     []byte("http://192.168.206.6/2016s/dk/top.asp?" + termInfo),
			ScoreUrl:    []byte("http://192.168.206.6/2016s/dk/bottom_list_new.asp?" + termInfo),
			SubscoreUrl: []byte("http://192.168.206.6/2016s/exam_a_p_g/bottom_list.asp?" + termInfo),
		}
	*/

	/*
	生态高三上
	termInfo := "year_in=2016&grade_name=%B8%DF%C8%FD&class_name1=2018-2019%C9%CF%D1%A7%C6%DA"
	return &FromWebsite{
		ExamUrl:     []byte("http://192.168.206.6/20162018/dk/top.asp?" + termInfo),
		ScoreUrl:    []byte("http://192.168.206.6/20162018/dk/bottom_list_new.asp?" + termInfo),
		SubscoreUrl: []byte("http://192.168.206.6/20162018/exam_a_p_g/bottom_list.asp?" + termInfo),
	}
	*/

	/*
	东城高三上
	termInfo := "year_in=2016&grade_name=%B8%DF%C8%FD&class_name1=2018-2019%C9%CF%D1%A7%C6%DA"
	return &FromWebsite{
		ExamUrl:     []byte("http://192.168.1.88/2016abc/dk/top.asp?" + termInfo),
		ScoreUrl:    []byte("http://192.168.1.88/2016abc/dk/bottom_list_new.asp?" + termInfo),
		SubscoreUrl: []byte("http://192.168.1.88/2016abc/exam_a_p_g/bottom_list.asp?" + termInfo),
	}
	*/


	termInfo := "year_in=2016&grade_name=%B8%DF%C8%FD&class_name1=2018-2019%C9%CF%D1%A7%C6%DA"
	return &FromWebsite{
		ExamUrl:     []byte("http://192.168.206.6/20162018/dk/top.asp?" + termInfo),
		ScoreUrl:    []byte("http://192.168.206.6/20162018/dk/bottom_list_new.asp?" + termInfo),
		SubscoreUrl: []byte("http://192.168.206.6/20162018/exam_a_p_g/bottom_list.asp?" + termInfo),
	}
}
