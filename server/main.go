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
	http.ListenAndServe(":8080", nil)
}

func genHandler(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("t")
	fmt.Println("t", t)
	rs, err := findPage(t)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`%s`, err.Error())))
	}
	w.WriteHeader(200)
	w.Write(rs)
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

	var objMap map[string]json.RawMessage
	json.Unmarshal(b, &objMap)

	var queryMap map[string]json.RawMessage
	json.Unmarshal(objMap["query"], &queryMap) // one level down
	
	return queryMap["search"], nil
}