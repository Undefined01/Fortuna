/*
http://192.168.206.6/2016s/dk/top.asp?year_in=2016&grade_name=%B8%DF%B6%FE&class_name1=2017-2018%C9%CF%D1%A7%C6%DA

http://192.168.206.6/2016s/dk/bottom_list_new.asp?year_in=2016&grade_name=%B8%DF%B6%FE&class_name1=2017-2018%CF%C2%D1%A7%C6%DA
exam_storting=%D6%DC%B2%E2%D2%BB&km=%CE%EF%C0%ED&select_or=%CA%C7&class_name=09
*/
package main
import (
	"log"
	"fmt"
	
	"io"
	"os"
	"strconv"
	
	"net/http"
	"net/url"
	"io/ioutil"
	"regexp"

	"github.com/axgle/mahonia"
)

const (
	term_info = "year_in=2016&grade_name=%B8%DF%B6%FE&class_name1=2017-2018%CF%C2%D1%A7%C6%DA"
	exam_url = "http://192.168.206.6/2016s/dk/top.asp?" + term_info
	score_url = "http://192.168.206.6/2016s/dk/bottom_list_new.asp?" + term_info
)
var subject_list = [11][2]string{{"语文", "否"}, {"数学", "否"}, {"数学文", "否"}, {"数学理", "否"}, {"英语", "否"}, {"物理", "是"}, {"化学", "是"}, {"生物", "是"}, {"政治", "否"}, {"历史", "否"}, {"地理", "否"}}

var (
	gbk_encoder = mahonia.NewEncoder("gbk")
	gbk_decoder = mahonia.NewDecoder("gbk")
)

func main() {
	log.Println("正在获取考试列表……")
	rawdata := HttpGet(exam_url)
	rawdata = rawdata[2300:3300]
	rawdata = []byte(gbk_decoder.ConvertString(string(rawdata)))
	exam_list := GetExamList(rawdata)
	
	exam := SelectExam(exam_list)
	log.Println("选择了", exam)
	
	fout, err := os.Create("result.csv")
	if err != nil {
		log.Panic("打开结果文件失败")
	}
	defer fout.Close()
	io.WriteString(fout, gbk_encoder.ConvertString("科目"))

	var subjects = 0
	var score = make(map[string][]string)
	for _, subject := range subject_list {
		log.Println("正在获取", subject[0], "成绩")
		score_list := GetScoreList(score_url, exam, "09", subject[0], subject[1])
		if len(score_list) == 0 { continue }
		subjects++
		io.WriteString(fout, gbk_encoder.ConvertString(", "))
		io.WriteString(fout, gbk_encoder.ConvertString(subject[0]))
		log.Println("共有", len(score_list), "人")
		for _, person := range score_list {
			if len(score[string(person[0])]) == 0 {
				score[string(person[0])] = make([]string, subjects - 1)
			}
			score[string(person[0])] = append(score[string(person[0])], string(person[1]))
		}
	}
	
	io.WriteString(fout, gbk_encoder.ConvertString(", 总分"))
	for k, v := range score {
		io.WriteString(fout, gbk_encoder.ConvertString("\n"))
		io.WriteString(fout, gbk_encoder.ConvertString(k))
		sum := 0.0
		for _, s := range v {
			num, _ := strconv.ParseFloat(s, 32)
			sum = num + sum
			io.WriteString(fout, gbk_encoder.ConvertString(", "))
			io.WriteString(fout, gbk_encoder.ConvertString(s))
		}
		fmt.Fprintf(fout, ", %f", sum)
	}
	io.WriteString(fout, gbk_encoder.ConvertString(", "))
	fmt.Printf("%v\n", score)
}

func HttpGet(requrl string) []byte {
	res, err := http.Get(requrl)
	if err != nil {
		log.Panic("无法连接至服务器：", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic("无法读取响应数据：", err)
	}
	return body
}

func HttpPostP(requrl string, para url.Values) []byte {
	res,err := http.PostForm(requrl, para)
    if err != nil {
		log.Panic("无法连接至服务器：", err)
    }
    defer res.Body.Close()
	
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
		log.Panic("无法读取响应数据：", err)
    }
    return body
}

func GetExamList(data []byte) [][]byte {
	var exam_list [][]byte
	
	index := fastfind(data, 0, "</select>")
	clip := data[:index]
	
	reg := regexp.MustCompile("<option value=(.*?)>")
	list := reg.FindAll(clip, 20)
	for i, value := range list {
		if i == 0 { continue }
		exam_list = append(exam_list, value[14 : len(value)-2])
	}
	return exam_list
}

func SelectExam(list [][]byte) string {
	println("请输入需要获取的考试id：")
	for i, value := range list {
		println(i, ":", string(value))
	}
	var id int
	fmt.Scanf("%d", &id)
	if id < 0 || id >= len(list) {
		log.Panic("id范围错误！")
	}
	return string(list[id])
}

func GetScoreList(requrl string, exam string, class string, subject string, select_or string) [][][]byte {
	exam = gbk_encoder.ConvertString(exam)
	subject = gbk_encoder.ConvertString(subject)
	select_or = gbk_encoder.ConvertString(select_or)
	postdata := url.Values{
		"exam_storting": {exam},
		"km": {subject},
		"select_or": {select_or},
		"class_name": {class},
	}
	res := HttpPostP(requrl, postdata)
	res = res[3200:]
	res = []byte(gbk_decoder.ConvertString(string(res)))
	
	var score_list [][][]byte
	reg := regexp.MustCompile("<p align=\"center\">(.*?)</")
	l := fastfind(res, 0, "<tr")
	for l != -1 {
		r := fastfind(res, l, "</tr>")
		list := reg.FindAllSubmatch(res[l:r], 20)
		if len(list) == 0 { break }
		if list[0][1][0] == ' ' { break }
		score_list = append(score_list, [][]byte{list[1][1], list[9][1]})
		l = fastfind(res, r + 4, "<tr")
	}
	return score_list
}

func fastfind(str []byte, pos int, pattern string) int {
	strlen := len(str)
	patlen := len(pattern)
	i := pos; j := 0
	for i < strlen {
		if str[i] == pattern[j] { j++ } else { i = i + j; j = 0 }
		if j == patlen { return i - j + 1 }
		i++;
	}
	return -1
}