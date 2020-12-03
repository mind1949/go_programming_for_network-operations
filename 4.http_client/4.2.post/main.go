package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	u := "https://golang.org/pkg/net/http/"
	data := []byte(`"hello": "world"`)
	req, _ := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	cookie := http.Cookie{Name: "cookiename", Value: "cookievalue"}
	req.AddCookie(&cookie)
	client := http.Client{Timeout: time.Second * 10}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s\n", body)
}
