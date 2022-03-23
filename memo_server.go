package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

const saveFile = "memo.txt"

func main() {
	print("memo server - [URL] http://localhost:8888/\n")
	http.HandleFunc("/", readHandler)
	http.HandleFunc("/w", writHandler)
	http.ListenAndServe(":8888", nil)
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	text, err := ioutil.ReadFile(saveFile)
	if err != nil {
		text = []byte("memo here")
	}
	htmlText := html.EscapeString(string(text))
	s := "<html>" +
		"<style>textarea { width:99%; height:200px; }</style>" +
		"<form method='POST' action='/W'>" +
		"<textarea name='text'>" + htmlText + "</textarea>" +
		"<input type='submit' value='保存' /></form></html>"
	w.Write([]byte(s))
}

func writHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form["text"]) == 0 {
		w.Write([]byte("フォームから投稿してください。"))
		return
	}
	text := r.Form["text"][0]
	ioutil.WriteFile(saveFile, []byte(text), 0644)
	fmt.Println("save:" + text)
	http.Redirect(w, r, "/", 301)
}
