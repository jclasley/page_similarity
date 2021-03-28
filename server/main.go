package main

import (
	"fmt"

	// "log"
	"net/http"
	"controller"
)

func main() {
	fs := http.FileServer(http.Dir("../public/"))
	http.Handle("/", fs)

	http.HandleFunc("/api/search/", genHandler)
	http.HandleFunc("/api/extract/", extractHandler) 
	http.ListenAndServe(":8080", nil)
}

func genHandler(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("t")
	fmt.Println("t", t)
	js, err := controller.FindPage(t)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`%s`, err.Error())))
	}
	w.WriteHeader(200)
	w.Write(js)
}

func extractHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println(id)
	js, err := controller.PlainTextPage(id)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`%s`, err.Error())))
	}
	w.WriteHeader(200)
	w.Write(js)
}