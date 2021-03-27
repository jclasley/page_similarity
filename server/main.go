package main

import (
	"encoding/json"
	"fmt"
	"io"

	// "log"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/api/search/", genHandler)
	http.HandleFunc("/api/extract/", extractHandler) 
	http.ListenAndServe(":8080", nil)
}

func genHandler(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("t")
	fmt.Println("t", t)
	js, err := findPage(t)
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
	js, err := plainTextPage(id)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`%s`, err.Error())))
	}
	w.WriteHeader(200)
	w.Write(js)
}

func findPage(t string) (json.RawMessage, error) {
	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&list=search&srsearch=%s&utf8=&format=json",
		url.PathEscape(t))
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	var objMap map[string]map[string]json.RawMessage
	json.Unmarshal(b, &objMap)
	
	return objMap["query"]["search"], nil
}

func plainTextPage(id string) (json.RawMessage, error) {
	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&format=json&pageids=%s&prop=extracts&explaintext", id)
	res, err := http.Get(url)
	fmt.Println(url)
	if err != nil {
		return []byte{}, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	var js map[string]map[string]map[string]map[string]json.RawMessage
	json.Unmarshal(b, &js)

	return js["query"]["pages"][id]["extract"], nil
}

// https://en.wikipedia.org/w/api.php?action=query&pageids=417121&prop=extracts&explaintext
// 