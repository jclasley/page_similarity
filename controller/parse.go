package controller

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"encoding/json"
)

func CheckSimilarity(t1 string, t2 string) float64 {
	r := strings.NewReplacer(":", "")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		t1 = r.Replace(t1)
		c1 <- t1
	}()
	go func() {
		t2 = r.Replace(t2)
		c2 <- t2
	}()
	key := os.Getenv("DAND_KEY")
	u := "https://api.dandelion.eu/datatxt/sim/v1/?lang=en&text1=%s&text2=%s&token=%s"
	u = fmt.Sprintf(u, url.PathEscape(<-c1), url.PathEscape(<-c2), key)
	runes := []rune(u)
	if len(runes) > 4096 {
		u = string(runes[:4000]) + fmt.Sprintf("&token=%s", key)
	}
	
	res, err := http.Get(u)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	b, _ := io.ReadAll(res.Body)
	var sim map[string]float64
	json.Unmarshal(b, &sim)
	return float64(sim["similarity"])
}