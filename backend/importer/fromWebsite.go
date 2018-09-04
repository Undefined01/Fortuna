package importer

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/axgle/mahonia"

	"log"
	"strconv"

	"github.com/Undefined01/fortuna/backend/utils"
)

// 数据导入模块：从学校的成绩查询网站获取成绩。
type FromWebsite struct {
	// 学校成绩查询网站的地址，年级
	ExamUrl     []byte
	ScoreUrl    []byte
	SubscoreUrl []byte
}

// 内部辅助函数，在转换失败时返回-1
func (this *FromWebsite) toInt(str []byte) int {
	num, err := strconv.Atoi(string(str))
	if err != nil {
		return -1
	}
	return num
}
func (this *FromWebsite) toFloat(str []byte) float32 {
	num, err := strconv.ParseFloat(string(str), 32)
	if err != nil {
		return -1
	}
	return float32(num)
}

// 发送HTTP/GET请求，并解码为UTF8
func (this *FromWebsite) httpGet(requrl []byte) []byte {
	res, err := http.Get(string(requrl))
	if err != nil {
		log.Printf("无法连接到服务器`%s`", requrl)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("读取回应消息失败")
		return nil
	}

	gbk_decoder := mahonia.NewDecoder("gbk")
	return []byte(gbk_decoder.ConvertString(string(body)))
}

// 发送HTTP/POST请求，并解码为UTF8
func (this *FromWebsite) httpPost(requrl []byte, args *url.Values) []byte {
	gbk_encoder := mahonia.NewEncoder("gbk")
	for k := range *args {
		for v := range (*args)[k] {
			(*args)[k][v] = gbk_encoder.ConvertString((*args)[k][v])
		}
	}

	res, err := http.PostForm(string(requrl), *args)
	if err != nil {
		log.Printf("无法连接到服务器`%s`", requrl)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("读取回应消息失败")
		return nil
	}

	gbk_decoder := mahonia.NewDecoder("gbk")
	return []byte(gbk_decoder.ConvertString(string(body)))
}

// 特殊查找字符串函数，匹配失败时直接放弃该位置的匹配。
// 返回第一个匹配的首位置
func (this *FromWebsite) Find(src []byte, dst string, pos int) int {
	srcLen := len(src)
	dstLen := len(dst)
	for i, j := pos, 0; i < srcLen; i++ {
		if src[i] == dst[j] {
			j++
			if j == dstLen {
				return i - j + 1
			}
		} else {
			j = 0
		}
	}
	return -1
}

// 获取两个关键字中间的字符串
// 返回该字符串的首末位置 [start, end)
func (this *FromWebsite) Between(src []byte, dst1 string, dst2 string, pos int) [2]int {
	srcLen := len(src)
	dstLen := len(dst1)
	for pos < srcLen {
		lpos := this.Find(src, dst1, pos)
		if lpos == -1 {
			break
		}
		pos = lpos + dstLen
		rpos := this.Find(src, dst2, pos)
		if rpos != -1 {
			return [2]int{lpos + dstLen, rpos}
		}
	}
	return [2]int{-1, -1}
}

// 获取所有“两个关键字中间的字符串”
// 返回包含所有字符串的数组
func (this *FromWebsite) AllBetween(src []byte, dst1 string, dst2 string, pos int) [][]byte {
	res := make([][]byte, 0, 40)
	for {
		index := this.Between(src, dst1, dst2, pos)
		if index[0] == -1 {
			break
		}
		res = append(res, src[index[0]:index[1]])
		pos = index[1]
	}
	return res
}

// 获取考试列表；返回考试名称的数组。
func (this *FromWebsite) GetExamList() []string {
	rawData := this.httpGet(this.ExamUrl)
	if rawData == nil {
		return nil
	}

	index := this.Between(rawData, "selected", "</select>", 0)
	if index[0] == -1 {
		log.Printf("意料之外的数据格式：找不到定位符`<select`")
		return nil
	}
	rawData = rawData[index[0]:index[1]]

	pos := 0
	examList := make([]string, 0, 20)
	for {
		index := this.Between(rawData, "<option value=", " >", pos)
		if index[0] != -1 {
			examList = append(examList, string(rawData[index[0]:index[1]]))
			pos = index[1]
		} else {
			break
		}
	}
	log.Printf("获取到%d场考试", len(examList))
	return examList
}

// 获取某一科目的得分，需要带上“是否文理选科”
func (this *FromWebsite) GetScoreX(exam string, class string, subject string, selector string) []utils.ScoreOfSubject {
	postData := &url.Values{
		"exam_storting": {exam},
		"km":            {subject},
		"select_or":     {selector},
		"class_name":    {class},
	}
	rawData := this.httpPost(this.ScoreUrl, postData)
	if rawData == nil {
		return nil
	}

	// 如果没有数据，直接返回nil
	pos := this.Find(rawData, "没有数据", 0)
	if pos != -1 {
		log.Printf("于 #%d 找到定位符`没有数据`", pos)
		return nil
	}

	// 跳过第一行标题数据
	pos = this.Find(rawData, "</tr>", 0)
	if pos == -1 {
		log.Printf("意料之外的数据格式：找不到定位符`</tr>`")
		return nil
	}
	log.Printf("于 #%d 找到定位符`</tr>`", pos)
	pos = this.Find(rawData, "</tr>", pos+5)
	if pos == -1 {
		log.Printf("意料之外的数据格式：找不到定位符`</tr>`")
		return nil
	}
	log.Printf("于 #%d 找到定位符`</tr>`", pos)
	rawData = rawData[pos:]

	pos = 0
	scoreList := make([]utils.ScoreOfSubject, 0, 60)
	for {
		index := this.Between(rawData, "<tr", "</tr>", pos)
		if index[0] == -1 {
			break
		}
		pos = index[1]
		row := rawData[index[0]:index[1]]
		rowData := this.AllBetween(row, "<p align=\"center\">", "</", 0)
		if len(rowData) < 9 {
			continue
		}

		sid := this.toInt(rowData[0])
		name := string(rowData[1])
		object := this.toFloat(rowData[7])
		subject := this.toFloat(rowData[8])
		total := this.toFloat(rowData[9])
		rank := this.toInt(rowData[6])
		scoreList = append(scoreList, utils.ScoreOfSubject{
			sid, name, object, subject, total, rank,
		})
	}

	log.Printf("共获取到 %d 行数据", len(scoreList))
	return scoreList
}

// 获取某一行中的详细得分，辅助GetSubscoreX
func (this *FromWebsite) getSubscoreRow(rowStr []byte, colLen int) []interface{} {
	row := this.AllBetween(rowStr, "<p align=center>", "</", 0)
	if len(row) != colLen {
		log.Printf("意料之外的数据格式：表格长度 %d 不匹配\n原始数据：\n%q", len(row), rowStr)
		return nil
	}

	sid := this.toInt(row[0])
	name := string(row[1])
	rowData := make([]interface{}, 0, 40)
	rowData = append(rowData, sid, name)
	for _, v := range row[2:] {
		rowData = append(rowData, this.toFloat(v))
	}
	return rowData
}

// 获取某一科目的小题得分，需要带上“是否文理选科”
func (this *FromWebsite) GetSubscoreX(exam string, class string, subject string, selector string) *utils.TableData {
	postdata := &url.Values{
		"exam_storting": {exam},
		"km":            {subject},
		"select_or":     {selector},
		"class_name":    {class},
	}
	rawData := this.httpPost(this.SubscoreUrl, postdata)
	if rawData == nil {
		return nil
	}

	// 如果没有数据，直接返回nil
	pos := this.Find(rawData, "没有数据", 0)
	if pos != -1 && pos <= 5000{
		log.Printf("于 #%d 找到定位符`没有数据`", pos)
		return nil
	}

	pos = this.Find(rawData, "</table>", 0)
	if pos == -1 {
		log.Println("意料之外的数据格式：找不到定位符`</table>`")
		return nil
	}
	log.Printf("于 #%d 找到定位符`</table>`", pos)
	rawData = rawData[pos:]

	var subscore utils.TableData
	subscore.Title = subject

	// 获取表头
	index := this.Between(rawData, "<tr", "<tr>", 0)
	subscore.Cols = make([]string, 0, 40)
	colData := rawData[index[0]:index[1]]
	colsTemp := this.AllBetween(colData, "<p align=center>", "</", 0)
	for _, v := range colsTemp {
		subscore.Cols = append(subscore.Cols, string(v))
	}
	colLen := len(subscore.Cols)
	log.Printf("共获取到 %d 列标题", colLen)

	// 截取，忽略试题分析
	pos = this.Find(rawData, "</table>", index[1])
	if pos == -1 {
		log.Println("意料之外的数据格式：找不到定位符`</table>`")
		return nil
	}
	log.Printf("于 #%d 找到定位符`</table>`", pos)
	rawData = rawData[index[1]+4 : pos]

	// 获取一行数据
	subscore.Data = make([][]interface{}, 0, 60)
	list := this.AllBetween(rawData, "<tr", "<tr", 0)
	for _, v := range list {
		rowData := this.getSubscoreRow(v, colLen)
		if rowData != nil {
			subscore.Data = append(subscore.Data, rowData)
		}
	}

	// 处理最后一个人
	lastPerson := this.Between(rawData, string(list[len(list)-1]), "\n", 0)
	rowData := this.getSubscoreRow(rawData[lastPerson[0]:lastPerson[1]], colLen)
	if rowData != nil {
		subscore.Data = append(subscore.Data, rowData)
	}

	log.Printf("共获取到 %d 行数据", len(subscore.Data))
	return &subscore
}

func (this *FromWebsite) GetScore(exam string, class string, subject string) []utils.ScoreOfSubject {
	log.Printf("正在获取 %s %s班 %s 的得分", exam, class, subject)
	score := this.GetScoreX(exam, class, subject, "否")
	if score != nil {
		return score
	} else {
		return this.GetScoreX(exam, class, subject, "是")
	}
}

func (this *FromWebsite) GetSubscore(exam string, class string, subject string) *utils.TableData {
	log.Printf("正在获取 %s %s班 %s 的小题得分", exam, class, subject)
	subscore := this.GetSubscoreX(exam, class, subject, "否")
	if subscore != nil {
		return subscore
	} else {
		return this.GetSubscoreX(exam, class, subject, "是")
	}
}
