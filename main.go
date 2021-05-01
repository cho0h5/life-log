package main

import (
	"fmt"
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
	db = initDB(DB_DSN)

	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)

	fmt.Printf("server start: localhost:%d\n", 8080)
	http.ListenAndServe(":8080", nil)
}
