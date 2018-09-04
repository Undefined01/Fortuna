package main

import (
	"regexp"
)

const (
	term_info = "year_in=2016&grade_name=%B8%DF%B6%FE&class_name1=2017-2018%CF%C2%D1%A7%C6%DA"
	exam_url = "http://192.168.206.6/2016s/dk/top.asp?" + term_info
	score_url = "http://192.168.206.6/2016s/dk/bottom_list_new.asp?" + term_info
	subscore_url = "http://192.168.206.6/2016s/exam_a_p_g/bottom_list.asp?" + term_info
)

var subject_list = map[string]string{
	"语文": "否",
	"数学": "否",
	"数学文": "否",
	"数学理": "否",
	"英语": "否",
	"物理": "是",
	"化学": "是",
	"生物": "是",
	"政治": "否",
	"历史": "否",
	"地理": "否",
}

func GetExamList() *[]string {
	data := HttpGet(exam_url)
	data = data[2448:]
	
	{
		re := regexp.MustCompile("</select>")
		index := re.FindIndex(data)
		data = data[:index[0]]
	}
	
	{
		var exam_list []string
		reg := regexp.MustCompile("<option value=(.*?)>")
		list := reg.FindAll(data, 20)
		for _, value := range list {
			exam_list = append(exam_list, string(value[14 : len(value)-2]))
		}
		return &exam_list
	}
}

func GetScoreList(exam string, class string, subject string) *map[string][3]float32 {
	select_or, ok := subject_list[subject]
	if !ok {
		return nil
	}
	postdata := map[string][]string{
		"exam_storting": {exam},
		"km": {subject},
		"select_or": {select_or},
		"class_name": {class},
	}
	data := HttpPost(score_url, postdata)
	data = data[3379:]
	
	re := regexp.MustCompile("(?ms)<tr.*?</tr>")
	list := re.FindAll(data, -1)
	if len(list) <= 3 { return nil }
	list = list[:len(list)-3]
	
	score_list := make(map[string][3]float32)
	re = regexp.MustCompile("<p align=\"center\">(.*?)</")
	for _, row := range list {
		row_data := re.FindAll(row, -1)
		name := row_data[1][18 : len(row_data[1])-2]
		object := row_data[7][18 : len(row_data[7])-2]
		subject := row_data[8][18 : len(row_data[8])-2]
		total := row_data[9][18 : len(row_data[9])-2]
		score_list[string(name)] = [3]float32{
			BytesToFloat(object),
			BytesToFloat(subject),
			BytesToFloat(total),
		}
	}
	return &score_list
}

func GetSubscore(exam string, class string, subject string) *map[string]map[int]float32 {
	select_or, ok := subject_list[subject]
	if !ok {
		return nil
	}
	postdata := map[string][]string{
		"exam_storting": {exam},
		"km": {subject},
		"select_or": {select_or},
		"class_name": {class},
	}
	data := HttpPost(subscore_url, postdata)
	data = data[2380:]
	
	re := regexp.MustCompile("<tr>")
	index := re.FindIndex(data)
	
	question_list := make([]int, 0, 30)
	{
		clip := data[:index[0]]
		re := regexp.MustCompile("<p align=center>(.*?)</")
		list := re.FindAll(clip, -1)
		for _, v := range list {
			question_list = append(question_list,
				BytesToInt(v[16 : len(v)-2]))
		}
	}
	
	subscore := make(map[string]map[int]float32)
	data = data[index[1]:]
	re = regexp.MustCompile("</table>")
	index = re.FindIndex(data)
	data = data[:index[0]-27]
	{
		re := regexp.MustCompile("<tr>(.*?)(<tr|\n)")
		list1 := re.FindAll(data, -1)
		re = regexp.MustCompile("<tr  class='alt'>(.*?)(<tr|\n)")
		list2 := re.FindAll(data, -1)
		
		list := make([][]byte, len(list1) + len(list2))
		copy(list, list1)
		copy(list[len(list1):], list2)
		
		re = regexp.MustCompile("<p align=center>(.*?)</")
		for _, v := range list {
			row := re.FindAll(v, -1)
			name := string(row[1][16 : len(row[1])-2])
			row = row[2:]
			score := make(map[int]float32)
			for i, v := range row {
				score[question_list[i]] = BytesToFloat(v[16 : len(v)-2])
			}
			subscore[name] = score
		}
	}
	
	return &subscore
}
