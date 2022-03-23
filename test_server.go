package main

import "net/http"

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/aaa", aaaHandler)
	http.HandleFunc("/bbb", bbbHandler)

	http.ListenAndServe(":8888", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func aaaHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa"))
}

func bbbHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bbb"))
}
