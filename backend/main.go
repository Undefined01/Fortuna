package main

import (
	"encoding/json"
	"github.com/Undefined01/fortuna/backend/importer"
	"github.com/Undefined01/fortuna/backend/summary"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	http.HandleFunc("/api", apiResponse)
	http.Handle("/", http.FileServer(http.Dir("./")))

	log.Println("正在监听8081端口")
	http.ListenAndServe(":8081", nil)
	log.Println("HTTP服务已停止")
}

func apiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")

	urlPrefix := "http://192.168.206.6/20162018"
	termInfo := "year_in=2016&grade_name=%B8%DF%C8%FD&class_name1=2018-2019%C9%CF%D1%A7%C6%DA"
	getter := &importer.FromWebsite{
		ExamUrl:     []byte(urlPrefix + "/dk/top.asp?" + termInfo),
		ScoreUrl:    []byte(urlPrefix + "/dk/bottom_list_new.asp?" + termInfo),
		SubscoreUrl: []byte(urlPrefix + "/exam_a_p_g/bottom_list.asp?" + termInfo),
	}
	apitype := r.FormValue("type")
	switch apitype {
	case "examlist":
		exam := getter.GetExamList()
		data, err := json.Marshal(exam)
		if err != nil {
			w.Write([]byte("{}"))
		} else {
			w.Write(data)
		}
	case "examdata":
		exam := r.FormValue("exam")
		class := r.FormValue("class")
		subscore := GetExamData(exam, class)
		data, err := json.Marshal(subscore)
		if err != nil {
			w.Write([]byte("{}"))
		} else {
			w.Write(data)
		}
	}
}

func GetExamData(exam string, class string) []importer.Table {
	return summary.GetExamData(exam, class)
}
