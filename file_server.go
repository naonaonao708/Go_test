package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("pub"))
	http.Handle("/", fs)
	http.HandleFunc("/dice", DiceHandler)
	http.ListenAndServe(":8888", nil)
}

func DiceHandler()
