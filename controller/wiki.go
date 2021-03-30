package controller

import (
	"fmt"
	"io"
	"net/url"
	"net/http"
	"encoding/json"
)

func FindPage(t string) (json.RawMessage, error) {
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

func PlainTextPage(id string) (json.RawMessage, error) {
// https://en.wikipedia.org/w/api.php?action=query&pageids=417121&prop=extracts&explaintext

	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&format=json&pageids=%s&prop=extracts&exintro&explaintext", id)
	res, err := http.Get(url)
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