package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var headers = map[string]string{
	"Referer":      "https://ting55.com/",
	"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36}",
	"Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
}

func get(url string, path string) error {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	for k, v := range headers {
		reqest.Header.Add(k, v)
	}
	var response *http.Response
	response, err = client.Do(reqest)
	if err != nil {
		return err
	}
	fmt.Println(response.Status)
	defer response.Body.Close()
	file, _ := os.Create(path)
	defer file.Close()
	io.Copy(file, response.Body)
	return nil
}

func main() {
	err := get("https://ting55.com/book/387-6", `F:/ximavip/6.map3`)
	fmt.Println(err)
}
