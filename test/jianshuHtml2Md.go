package test

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var reFormat = regexp.MustCompile(`\{(\d*)\}`)

func formatStr(str string, args ...string) string {
	i := -1
	return reFormat.ReplaceAllStringFunc(str, func(s string) string {
		s = strings.Trim(s, "{}")
		i++
		if s == "" {
			return args[i]
		} else {
			index, _ := strconv.ParseInt(s, 10, 32)
			return args[index]
		}
	})
}

func bytesMd5(b []byte) string {
	has := md5.New()
	has.Write(b)
	return fmt.Sprintf("%x", has.Sum(nil))
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func downloadImg(url string, imgPath string) (string, error) {
	if exist, _ := pathExists(imgPath); !exist {
		os.MkdirAll(imgPath, 0644)
	}
	suffix := ".png"
	lastIndex := strings.LastIndex(url, ".")
	if lastIndex > 0 {
		suffix_ := url[lastIndex:len(url)]
		if suffix_ == ".jpg" || suffix_ == ".png" || suffix_ == ".gif" {
			suffix = suffix_
		}
	}
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}
	body, _ := io.ReadAll(res.Body)
	name := bytesMd5(body) + suffix
	imgPath = path.Join(imgPath, name)
	if exist, _ := pathExists(imgPath); !exist {
		file, _ := os.Create(imgPath)
		defer file.Close()
		io.Copy(file, bytes.NewBuffer(body))
	}
	return name, nil
}

func markdownUseLocalImg(text string, imgPath string) string {
	re := regexp.MustCompile(`!\[\]\((http.*?)\)`)
	subStrs := re.FindAllStringSubmatch(text, -1)
	for _, strs := range subStrs {
		url := strs[1]
		name, err := downloadImg(url, imgPath)
		if err == nil {
			text = strings.ReplaceAll(text, url, "assets/"+name)
		} else {
			fmt.Printf("下载失败:%s\n", text)
			fmt.Println(err)
		}
	}
	return text
}

func replaceWithDiv(selection *goquery.Selection, text string) {
	node := &html.Node{
		Type:     html.ElementNode,
		Data:     "div",
		DataAtom: atom.Div,
	}
	node.AppendChild(&html.Node{
		Type: html.TextNode,
		Data: text,
	})
	selection.ReplaceWithNodes(node)
}

func convertLink(doc *goquery.Selection) {
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		href = strings.TrimSpace(href)
		text := selection.Text()
		text = strings.TrimSpace(text)
		fmt.Println(href, text)
		if href != "" || text != "" {
			selection.ReplaceWithHtml(formatStr(`[{1}]({0} "{1}")`, href, text))
			//selection.Parent().SetText(formatStr(`[{1}]({0} "{1}")`, href, text))
		}
	})
}

func convertImg(doc *goquery.Selection) {
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("src")
		if strings.HasPrefix(url, "//") {
			url = "https:" + url
		}
		selection.ReplaceWithHtml(formatStr(` ![]({0}) `, url))
		//selection.Parent().SetText(formatStr(` ![]({0}) `, url))
	})
}

func convertHtag(doc *goquery.Selection) {
	titlePrefix := "##"
	hs := doc.Find("h1,h2,h3,h4,h5,h6,h7,h8,h9")
	var nums []int
	hs.Each(func(i int, selection *goquery.Selection) {
		num, _ := strconv.ParseInt(strings.Trim(selection.Nodes[0].Data, "h"), 10, 32)
		nums = append(nums, int(num))
	})
	var hLevel = make(map[int]int)
	sort.Ints(nums)
	var maxLevel int
	for _, num := range nums {
		level, exist := hLevel[num]
		if !exist {
			hLevel[num] = maxLevel
			maxLevel++
		} else {
			hLevel[num] = level
		}
	}
	hs.Each(func(i int, selection *goquery.Selection) {
		num, _ := strconv.ParseInt(strings.Trim(selection.Nodes[0].Data, "h"), 10, 32)
		text := selection.Text()
		text = formatStr(`{}{} {}`, titlePrefix, strings.Repeat("#", hLevel[int(num)]), text)
		//selection.ReplaceWithHtml(text)
		replaceWithDiv(selection, text)
	})
}

func convertBtag(doc *goquery.Selection) {
	doc.Find("b,strong").Each(func(i int, selection *goquery.Selection) {
		node := selection.Nodes[0]
		node.InsertBefore(&html.Node{
			Type: html.TextNode,
			Data: "**",
		}, node.FirstChild)
		node.AppendChild(&html.Node{
			Type: html.TextNode,
			Data: "**",
		})
	})
}

func convertUl(doc *goquery.Selection) {
	doc.Find("ul,ol").Each(func(i int, selection *goquery.Selection) {
		texts := selection.ChildrenFiltered("li").Map(func(i int, li *goquery.Selection) string {
			return "* " + strings.TrimSpace(li.Text())
		})
		text := strings.Join(texts, "    \n")
		replaceWithDiv(selection, text)
	})
}

func convertCode(doc *goquery.Selection) {
	doc.Find("code,pre").Each(func(i int, selection *goquery.Selection) {
		code := selection.Text()
		var text string
		if strings.Contains(code, "\n") {
			text = fmt.Sprintf("```%s\n%s\n```", "py", code)
		} else {
			text = fmt.Sprintf("`%s`", code)
		}
		replaceWithDiv(selection, text)
	})
}

func convertBlockquote(doc *goquery.Selection) {
	doc.Find("blockquote").Each(func(i int, selection *goquery.Selection) {
		code := selection.Text()
		text := fmt.Sprintf("> %s\n\n", code)
		replaceWithDiv(selection, text)
	})
}

func convertTable(doc *goquery.Selection) {
	doc.Find("table").Each(func(i int, selection *goquery.Selection) {
		var txts []string
		selection.Find("tr").Each(func(i int, selection *goquery.Selection) {
			texts := selection.ChildrenFiltered("td,th").Map(func(i int, li *goquery.Selection) string {
				return strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(li.Text()), "|", "&#124;"), "â¦", "...")
			})
			text := "| " + strings.Join(texts, " | ") + " |"
			txts = append(txts, text)
			if i == 0 {
				texts = selection.ChildrenFiltered("td,th").Map(func(i int, li *goquery.Selection) string {
					return "--"
				})
				txts = append(txts, "| "+strings.Join(texts, " | ")+" |")
			}
		})
		text := strings.Join(txts, "\n")
		replaceWithDiv(selection, text)
	})
}

func outTagText(node *html.Node, sep string, strip bool) string {
	lineAppend := "\n"
	if node.Type == html.TextNode {
		if strip {
			return strings.TrimSpace(node.Data)
		} else {
			return node.Data
		}
	}

	var texts []string
	var preChild *html.Node
	lineMap := map[string]int{"br": 1, "p": 1, "div": 1, "ul": 1, "ol": 1, "h1": 1, "h2": 1, "h3": 1, "h4": 1, "h5": 1}
	lineMap2 := map[string]int{"p": 1, "div": 1, "ul": 1, "ol": 1, "h1": 1, "h2": 1, "h3": 1, "h4": 1, "h5": 1}
	child := node.FirstChild
	for child != nil {
		if (preChild != nil && lineMap[preChild.Data] == 1) || lineMap2[child.Data] == 1 {
			length := len(texts)
			start := length - 4
			if start < 0 {
				start = 0
			}
			if strings.HasSuffix(strings.Join(texts[start:length], ""), "\n\n") && strings.Contains(lineAppend, "\n") {
				if !strings.HasSuffix(strings.Join(texts[start:length], ""), "\n\n\n") {
					texts = append(texts, "\n")
				}
			} else {
				texts = append(texts, lineAppend)
			}
		}
		preChild = child
		if child.Data == "script" || child.Data == "style" {
			continue
		}
		if child.Type == html.ElementNode {
			texts = append(texts, outTagText(child, sep, strip))
		} else if child.Type == html.TextNode {
			if strip {
				texts = append(texts, strings.TrimSpace(child.Data))
			} else {
				texts = append(texts, child.Data)
			}
		}
		child = child.NextSibling
	}
	return strings.Join(texts, sep)
}

// go get -u github.com/PuerkitoBio/goquery
func main2() {
	data, err := os.ReadFile("D:\\pycharmWork\\pyutils\\html2md\\test\\txt")
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	contentDiv := doc.Find("article")

	convertLink(contentDiv)
	convertImg(contentDiv)
	convertHtag(contentDiv)
	convertBtag(contentDiv)
	convertUl(contentDiv)
	convertCode(contentDiv)
	convertBlockquote(contentDiv)
	convertTable(contentDiv)

	fmt.Println(strings.Repeat("-", 60) + "\n")
	text := outTagText(contentDiv.Nodes[0], "", true)
	text = markdownUseLocalImg(text, "assets")
	fmt.Println(text)
	fmt.Println(strings.Repeat("-", 60) + "\n")
	//fmt.Println(contentDiv.Text())
	//fmt.Println(strings.Repeat("-", 60))
	//fmt.Println(contentDiv.Html())
}
