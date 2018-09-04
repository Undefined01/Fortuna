package main

import (
	"log"
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/api", ApiResponse)
	http.Handle("/", http.FileServer(http.Dir("./")))

	log.Println("Listening...")
	http.ListenAndServe(":8081", nil)
}

func ApiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	

	apitype := r.FormValue("type")
	switch apitype {
	case "exam":
		exam := GetExamList()
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
		score := GetScoreList(exam, class, subject)
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
		subscore := GetSubscore(exam, class, subject)
		data, err := json.Marshal(subscore)
		if err != nil {
			w.Write([]byte("{}"))
		} else {
			w.Write(data)
		}
	}
	
}