package main

import (
	"fmt"

	// "log"
	"net/http"
	"controller"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("../public/"))
	http.Handle("/", fs)

	http.HandleFunc("/api/search/", genHandler)
	http.HandleFunc("/api/extract/", extractHandler)
	http.HandleFunc("/api/compare/", compareHandler)
	http.ListenAndServe(":8080", nil)
}

func genHandler(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("t")
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
	js, err := controller.PlainTextPage(id)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`%s`, err.Error())))
	}
	w.WriteHeader(200)
	w.Write(js)
}

func compareHandler(w http.ResponseWriter, r *http.Request) {
	t1 := r.URL.Query().Get("t1")
	t2 := r.URL.Query().Get("t2")
	os.Setenv("DAND_KEY", "d5fae78897c047fab1effd387ab9a5c8")
	s := controller.CheckSimilarity(t1, t2)
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf(`%v`, s)))
}