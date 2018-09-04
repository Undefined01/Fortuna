package server

import (
	"encoding/json"
	"github.com/Undefined01/fortuna/backend/importer"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/api", apiResponse)
	http.Handle("/", http.FileServer(http.Dir("./")))

	log.Println("正在监听8081端口")
	http.ListenAndServe(":8081", nil)
	log.Println("HTTP服务已停止")
}

func apiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")

	getter := importer.NewFromWebsiteImporter()
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
	case "score":
		exam := r.FormValue("exam")
		class := r.FormValue("class")
		subject := r.FormValue("subject")
		score := getter.GetScore(exam, class, subject)
		data, err := json.Marshal(score)
		if err != nil {
			w.Write([]byte("{}"))
		} else {
			w.Write(data)
		}
	case "subscore":
		exam := r.FormValue("exam")
		class := r.FormValue("class")
		subject := r.FormValue("subject")
		subscore := getter.GetSubscore(exam, class, subject)
		data, err := json.Marshal(subscore)
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
