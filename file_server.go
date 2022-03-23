package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("pub"))
	http.Handle("/", fs)
	http.HandleFunc("/dice", DiceHandler)
	http.ListenAndServe(":8888", nil)
}

func DiceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf(
		"<h1>サイコロ: %d</h1>", rand.Intn(6)+1)))
}
