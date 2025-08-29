package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Certificate struct {
	Username string `json:"username"`
	Date     string `json:"date"`
	Message  string `json:"message"`
}

func certificateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POSTで送信してください", http.StatusMethodNotAllowed)
		return
	}

	// JSONのリクエストを受け取る
	var input struct {
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// データを作る
	cert := Certificate{
		Username: input.Username,
		Date:     time.Now().Format("2006-01-02"), // 今日の日付
		Message:  "認定証を発行しました",
	}

	// JSONに変換して返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cert); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	 fs := http.FileServer(http.Dir("./public"))
    http.Handle("/", fs)
	http.HandleFunc("/certificate", certificateHandler)
	http.ListenAndServe(":8080", nil)
}
