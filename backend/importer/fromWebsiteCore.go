package importer

import (
	"log"
	"net/url"
)

// 从学校的成绩查询网站获取成绩的内核
type FromWebsiteCore struct {
	// 学校成绩查询网站的地址（包括年级信息）
	ExamUrl     []byte
	ScoreUrl    []byte
	SubscoreUrl []byte
}

// 解析考试列表；返回考试名称的数组。
func (this *FromWebsiteCore) ParseExamList(rawData []byte) []string {
	index := between(rawData, "selected", "</select>", 0)
	if index[0] == -1 {
		log.Printf("意料之外的数据格式：找不到定位符`<select`")
		return nil
	}
	rawData = rawData[index[0]:index[1]]

	pos := 0
	examList := make([]string, 0, 20)
	for {
		index := between(rawData, "<option value=", " >", pos)
		if index[0] != -1 {
			examList = append(examList, string(rawData[index[0]:index[1]]))
			pos = index[1]
		} else {
			break
		}
	}
	return examList
}

// 解析某一科目的得分
func (this *FromWebsiteCore) ParseScore(rawData []byte) ScoreList {
	// 如果没有数据，直接返回nil
	pos := find(rawData, "没有数据", 0)
	if pos != -1 {
		return nil
	}

	// 跳过第一行标题数据
	pos = find(rawData, "</tr>", 0)
	if pos == -1 {
		log.Printf("意料之外的数据格式：找不到定位符`</tr>`")
		return nil
	}
	pos = find(rawData, "</tr>", pos+5)
	if pos == -1 {
		log.Printf("意料之外的数据格式：找不到定位符`</tr>`")
		return nil
	}
	rawData = rawData[pos:]

	pos = 0
	scoreList := make(ScoreList, 0, 60)
	for {
		index := between(rawData, "<tr", "</tr>", pos)
		if index[0] == -1 {
			break
		}
		pos = index[1]
		row := rawData[index[0]:index[1]]
		rowData := allBetween(row, "<p align=\"center\">", "</", 0)
		if len(rowData) < 9 {
			continue
		}

		sid := toInt(rowData[0])
		name := string(rowData[1])
		object := toFloat(rowData[7])
		subject := toFloat(rowData[8])
		total := toFloat(rowData[9])
		rank := toInt(rowData[6])
		scoreList = append(scoreList, Score{
			sid, name, object, subject, total, rank,
		})
	}
	return scoreList
}

// 解析某一行中的详细得分，辅助 ParseSubscore
func (this *FromWebsiteCore) parseSubscoreRow(rowStr []byte, colLen int) (int, []interface{}) {
	row := allBetween(rowStr, "<p align=center>", "</", 0)
	if len(row) != colLen {
		log.Printf("意料之外的数据格式：表格长度 %d 不匹配\n原始数据：\n%q", len(row), rowStr)
		return -1, nil
	}

	sid := toInt(row[0])
	name := string(row[1])
	rowData := make([]interface{}, 0, 40)
	rowData = append(rowData, sid, name)
	for _, v := range row[2:] {
		rowData = append(rowData, toFloat(v))
	}
	return sid, rowData
}

// 解析某一科目的小题得分
func (this *FromWebsiteCore) ParseSubscore(rawData []byte) *RawSubscore {
	// 如果没有数据，直接返回nil
	pos := find(rawData, "没有数据", 0)
	if pos != -1 && pos <= 5000 {
		return nil
	}

	pos = find(rawData, "</table>", 0)
	if pos == -1 {
		log.Println("意料之外的数据格式：找不到定位符`</table>`")
		return nil
	}
	rawData = rawData[pos:]

	var subscore RawSubscore

	// 获取表头
	index := between(rawData, "<tr", "<tr>", 0)
	subscore.Cols = make([]string, 0, 40)
	colData := rawData[index[0]:index[1]]
	colsTemp := allBetween(colData, "<p align=center>", "</", 0)
	for _, v := range colsTemp {
		subscore.Cols = append(subscore.Cols, string(v))
	}
	colLen := len(subscore.Cols)

	// 截取，忽略试题分析
	pos = find(rawData, "</table>", index[1])
	if pos == -1 {
		log.Println("意料之外的数据格式：找不到定位符`</table>`")
		return nil
	}
	rawData = rawData[index[1]+4 : pos]

	// 获取一行数据
	subscore.Data = make(map[int][]interface{})
	list := allBetween(rawData, "<tr", "<tr", 0)
	for _, v := range list {
		sid, rowData := this.parseSubscoreRow(v, colLen)
		if rowData != nil {
			subscore.Data[sid] = rowData
		}
	}

	// 处理最后一个人
	lastPerson := between(rawData, string(list[len(list)-1]), "\n", 0)
	sid, rowData := this.parseSubscoreRow(rawData[lastPerson[0]:lastPerson[1]], colLen)
	if rowData != nil {
		subscore.Data[sid] = rowData
	}
	return &subscore
}

// 获取考试列表；返回考试名称的数组。
func (this *FromWebsiteCore) GetExamList() []string {
	rawData := httpGet(this.ExamUrl)
	if rawData == nil {
		return nil
	}
	return this.ParseExamList(rawData)
}

// 获取某一科目的得分，需要带上“是否文理选科”
func (this *FromWebsiteCore) GetScore(exam string, class string, subject string) ScoreList {
	postData := &url.Values{
		"exam_storting": {exam},
		"km":            {subject},
		"select_or":     {"否"},
		"class_name":    {class},
	}
	rawData := httpPost(this.ScoreUrl, postData)
	if rawData == nil {
		return nil
	}
	res := this.ParseScore(rawData)
	if res != nil {
		return res
	}

	postData = &url.Values{
		"exam_storting": {exam},
		"km":            {subject},
		"select_or":     {"是"},
		"class_name":    {class},
	}
	rawData = httpPost(this.ScoreUrl, postData)
	if rawData == nil {
		return nil
	}
	return this.ParseScore(rawData)
}

// 获取某一科目的小题得分
func (this *FromWebsiteCore) GetSubscore(exam string, class string, subject string) *RawSubscore {
	postdata := &url.Values{
		"exam_storting": {exam},
		"km":            {subject},
		"select_or":     {"否"},
		"class_name":    {class},
	}
	rawData := httpPost(this.SubscoreUrl, postdata)
	if rawData == nil {
		return nil
	}
	res := this.ParseSubscore(rawData)
	if res != nil {
		return res
	}

	postdata = &url.Values{
		"exam_storting": {exam},
		"km":            {subject},
		"select_or":     {"是"},
		"class_name":    {class},
	}
	rawData = httpPost(this.SubscoreUrl, postdata)
	if rawData == nil {
		return nil
	}
	return this.ParseSubscore(rawData)
}
