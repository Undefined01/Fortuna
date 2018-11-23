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

	http.HandleFunc("/api/examlist", apiExamList)
	http.HandleFunc("/api/examdata", apiExamData)
	http.Handle("/", http.FileServer(http.Dir("./")))

	log.Println("正在监听8081端口")
	http.ListenAndServe(":8081", nil)
	log.Println("HTTP服务已停止")
}

func apiExamList(w http.ResponseWriter, r *http.Request) {
	eco := importer.NewFromWebsite()
	exam := eco.GetExamList()
	data, err := json.Marshal(exam)
	if err != nil {
		w.Write([]byte("{}"))
	} else {
		w.Write(data)
	}
}

func apiExamData(w http.ResponseWriter, r *http.Request) {
	exam := r.FormValue("exam")
	class := r.FormValue("class")
	subscore := summary.GetExamData(exam, class)
	data, err := json.Marshal(subscore)
	if err != nil {
		w.Write([]byte("{}"))
	} else {
		w.Write(data)
	}
}
