package main

import (
	"html/template"
	"net/http"
	"time"
)

var db *DB

func index(w http.ResponseWriter, r *http.Request) {
	logs := []Log{}
	db.readLogs(&logs)

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, logs)
}

func create(w http.ResponseWriter, r *http.Request) {
	db.writeLog(r.FormValue(("log")), time.Now().Format(TIME_FORMAT))

	http.Redirect(w, r, "/", 301)
}

func main() {
	db = initDB(DB_FILE)

	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.ListenAndServe(":8080", nil)
}