package request

import (
	"encoding/json"
	"fmt"
	"github.com/mozillazg/request"
	"net/http"
	"testing"
)

func TestRequestGet(t *testing.T) {
	c := new(http.Client)
	req := request.NewRequest(c)
	resp, err := req.Get("http://127.0.0.1:8000/get?name=小明")
	if err != nil {
		fmt.Println(err)
	} else {
		content, _ := resp.Content()
		var ret map[string]interface{}
		json.Unmarshal(content, &ret)
		for k, v := range ret {
			fmt.Println(k, v)
		}
		bytes, _ := json.MarshalIndent(ret, "", "    ")
		fmt.Println(string(bytes))
	}
}

func TestRequestGet2(t *testing.T) {
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Params = map[string]string{"name2": "小花"}
	req.Data = map[string]string{"name3": "小亮"}
	resp, err := req.Get("http://127.0.0.1:8000/get?name=小明")
	if err != nil {
		fmt.Println(err)
	} else {
		content, _ := resp.Content()
		var ret map[string]interface{}
		json.Unmarshal(content, &ret)
		for k, v := range ret {
			fmt.Println(k, v)
		}
		bytes, _ := json.MarshalIndent(ret, "", "    ")
		fmt.Println(string(bytes))
	}
}

func TestRequestPost(t *testing.T) {
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Params = map[string]string{"name2": "小花"}
	req.Data = map[string]string{"name3": "小亮"}
	resp, err := req.Post("http://127.0.0.1:8000/post?name=小明")
	if err != nil {
		fmt.Println(err)
	} else {
		content, _ := resp.Content()
		var ret map[string]interface{}
		json.Unmarshal(content, &ret)
		for k, v := range ret {
			fmt.Println(k, v)
		}
		bytes, _ := json.MarshalIndent(ret, "", "    ")
		fmt.Println(string(bytes))
	}
}

func TestRequestPostJson(t *testing.T) {
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Params = map[string]string{"name2": "小花"}
	req.Json = map[string]string{"name3": "小亮"}
	resp, err := req.Post("http://127.0.0.1:8000/post?name=小明")
	if err != nil {
		fmt.Println(err)
	} else {
		content, _ := resp.Content()
		var ret map[string]interface{}
		json.Unmarshal(content, &ret)
		for k, v := range ret {
			fmt.Println(k, v)
		}
		bytes, _ := json.MarshalIndent(ret, "", "    ")
		fmt.Println(string(bytes))
	}
}
