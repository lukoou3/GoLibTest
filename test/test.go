package test

import (
	"github.com/PuerkitoBio/goquery"
)
import "strings"
import "fmt"

func main1() {
	var html = `
<div class="contson" id="contsonba4626c44270">
    <p>胜日寻芳泗水滨，无边光景一时新。</p>
    <p>等闲识得东风面，万紫千红总是春。</p>
</div>
<div class="contson" id="contson846e626d74d3">
    <p>墙角数枝梅，凌寒独自开。</p>
    <p>遥知不是雪，为有暗香来。</p>
</div>
`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err)
		return
	}

	doc.Find("p").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	doc.FindNodes()

}
