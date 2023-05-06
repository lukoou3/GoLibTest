package template

import (
	"fmt"
	"os"
	"testing"
	"text/template"
)

func testTemplate(text string, data interface{}) {
	// 初始化，解析
	tmpl, err := template.New("test").Parse(text)
	if err != nil {
		panic(err)
	}
	// 输出到 os.Stdout
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

type People struct {
	Name string
	Age  int
}

func (p *People) ToStr() string {
	return fmt.Sprintf("name:%s,age:%d.", p.Name, p.Age)
}
func TestStructField(t *testing.T) {
	data := People{Name: "小明", Age: 20}
	text := `name is {{.Name}}, age is {{.Age}}.{{.ToStr}}` + "\n"
	//testTemplate(text, data)
	testTemplate(text, &data) //值和指针都可以。但是方法必须匹配指针还是值，编译器没法自动转换。
}

func TestMapField(t *testing.T) {
	data := map[string]interface{}{
		"name": "小明",
		"age":  20,
	}
	text := `name is {{.name}}, age is {{.age}}` + "\n"
	testTemplate(text, data)
	testTemplate(text, &data) //值和指针都可以
}

func TestValue(t *testing.T) {
	data := 2
	text := `value is {{.}}` + "\n"
	testTemplate(text, data)
	testTemplate(text, &data) //值和指针都可以
}

func TestValue2(t *testing.T) {
	data := map[string]interface{}{
		"name": "小明",
		"age":  20,
	}
	text := `value is {{.}}` + "\n"
	testTemplate(text, data)
	testTemplate(text, &data) //值和指针都可以
}

func TestTrim(t *testing.T) {
	data := map[string]interface{}{
		"name1": "    aa   ",
		"name2": "小明",
		"name3": "   cc   ",
	}
	text := `{{.name1}}     {{ .name2 }}    {{.name3}}` + "\n"
	testTemplate(text, data)
	// Trim不会包含其他{{}}生成的值
	text = `{{.name1}}     {{- .name2 -}}    {{.name3}}` + "\n"
	testTemplate(text, data)
}

func TestFunc(t *testing.T) {
	text := `name is {{.name}}, func1 return:{{func1}}, func2 return:{{func2 .name}}` + "\n"
	// 初始化
	tmpl := template.New("test")
	var funcMap = template.FuncMap{
		"func1": func() string {
			return "func1"
		},
		"func2": func(a string) string {
			return "func2-" + a
		},
	}
	tmpl.Funcs(funcMap)
	tmpl, err := tmpl.Parse(text)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"name": "小明",
		"age":  20,
	}

	// 输出到 os.Stdout
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func TestIf(t *testing.T) {
	data := map[string]interface{}{
		"name": "小明",
		"age":  20,
	}
	text := `{{if .name}} name is {{.name}} age is {{.age}} {{else}} name is empty {{end}}` + "\n"
	testTemplate(text, data)

	data = map[string]interface{}{
		"name": "",
		"age":  20,
	}
	testTemplate(text, data)

	fmt.Println("---------------------")

	data = map[string]interface{}{
		"name": "小明",
		"age":  20,
	}
	text = `{{if ge .age  12}} name is {{.name}} age is {{.age}} {{else}} age < 12 {{end}}` + "\n"
	testTemplate(text, data)

	data = map[string]interface{}{
		"name": "小明",
		"age":  8,
	}
	text = `{{if ge .age  12}} name is {{.name}} age is {{.age}} {{else}} age < 12 {{end}}` + "\n"
	testTemplate(text, data)
}

func TestRange(t *testing.T) {
	data := map[string]interface{}{
		"names": []string{"小明", "小花", "小亮"},
	}
	text := `{{range .names}}{{.}}, {{end}}` + "\n"
	testTemplate(text, data)
	text = `{{range $i,$v := .names}}{{if gt $i  0}}, {{end}}{{$i}} -> {{$v}}{{end}}` + "\n"
	testTemplate(text, data)
	data = map[string]interface{}{
		"name": "小明",
		"age":  8,
	}
	text = `{{range $k,$v := .}}, {{$k}} -> {{$v}}{{end}}` + "\n"
	testTemplate(text, data)
}
