package test

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"
)

func http_get(url string) ([]byte, error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36")

	if err != nil {
		return nil, err
	}

	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	return body, nil
}

func getItemInfo(url string) string {
	body, _ := http_get(url)
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	text := doc.Find("#sonsyuanwen .contson").Text()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
	fmt.Println(text)
	return text
}

func urlJoin(baseUrl string, url2 string) string {
	base, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	ref, err := url.Parse(url2)
	if err != nil {
		log.Fatal(err)
	}
	u := base.ResolveReference(ref)
	return u.String()
}

func Test11(t *testing.T) {
	fmt.Println(url.JoinPath("https://so.gushiwen.cn/shiwens/", "/shiwenv_e8c06d3a532d.aspx"))
	base, err := url.Parse("https://so.gushiwen.cn/shiwens/")
	if err != nil {
		log.Fatal(err)
	}
	ref, err := url.Parse("/shiwenv_e8c06d3a532d.aspx")
	if err != nil {
		log.Fatal(err)
	}
	u := base.ResolveReference(ref)
	fmt.Println(u.String()) // prints http://foo/bar.html

}

func TestGushiwen(t *testing.T) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
	body, _ := http_get("https://so.gushiwen.cn/shiwens/")
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
	divs := doc.Find("#leftZhankai div.sons")

	results := make(chan string)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(divs.Nodes))

	divs.Each(func(i int, div *goquery.Selection) {
		href, _ := div.Find("div.cont > p > a").Attr("href")
		href2 := urlJoin("https://so.gushiwen.cn/shiwens/", href)
		println(href, href2)
		go func() {
			results <- getItemInfo(href2)
			waitGroup.Done()
		}()
	})

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
}
