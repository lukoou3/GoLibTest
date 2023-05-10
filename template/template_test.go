package template

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"text/template"
	"time"
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

func TestFuncSequence(t *testing.T) {
	// 初始化
	tmpl := template.New("test")
	var funcMap = template.FuncMap{
		"sequence": func(_start interface{}, _stop interface{}, _step interface{}) (rst []interface{}, err error) {
			switch _start.(type) {
			case int:
				if reflect.TypeOf(_stop).Kind() != reflect.Int || reflect.TypeOf(_step).Kind() != reflect.Int {
					err = errors.New("必须都是int类型")
				} else {
					start := _start.(int)
					stop := _stop.(int)
					step := _step.(int)
					if !((step > 0 && start <= stop) || (step < 0 && start >= stop) || (step == 0 && start == stop)) {
						err = errors.New("int大小不正确")
						return
					}
					size := 1
					if start != stop {
						size = (stop-start)/step + 1
					}
					rst = make([]interface{}, size)
					for i := 0; i < size; i++ {
						rst[i] = start + i*step
					}
				}
			case string:
				if len(_start.(string)) == 19 {
					var start, stop time.Time
					start, err = time.ParseInLocation(time.DateTime, _start.(string), time.Local)
					stop, err = time.ParseInLocation(time.DateTime, _stop.(string), time.Local)

					var step time.Duration
					step, err = time.ParseDuration(_step.(string))
					if !((step > 0 && start.Compare(stop) <= 1) || (step < 0 && start.Compare(stop) >= 1) || (step == 0 && start.Compare(stop) == 0)) {
						err = errors.New("date大小不正确")
						return
					}
					size := 1
					if start.Compare(stop) != 0 {
						size = (int(stop.Sub(start)))/int(step) + 1
					}
					rst = make([]interface{}, size)
					var steps time.Duration = 0
					for i := 0; i < size; i++ {
						rst[i] = start.Add(steps).Format(time.DateTime)
						steps += step
					}
				} else {
					var start, stop time.Time
					start, err = time.ParseInLocation(time.DateOnly, _start.(string), time.Local)
					stop, err = time.ParseInLocation(time.DateOnly, _stop.(string), time.Local)

					var stepType = 1
					__step := _step.(string)
					if strings.HasSuffix(strings.ToLower(__step), "m") {
						stepType = 2
					}
					step, e := strconv.Atoi(__step[:len(__step)-1])
					if e != nil {
						e = err
						return
					}
					if !((step > 0 && start.Compare(stop) <= 1) || (step < 0 && start.Compare(stop) >= 1) || (step == 0 && start.Compare(stop) == 0)) {
						err = errors.New("date大小不正确")
						return
					}
					maxSize := 1
					if start.Compare(stop) != 0 {
						if stepType == 2 {
							maxSize = (int(stop.Sub(start)))/int(time.Hour*24*30) + 1
						} else {
							maxSize = (int(stop.Sub(start)))/int(time.Hour*24) + 1
						}

					}
					rst = make([]interface{}, maxSize)
					i := 0
					var date time.Time
					for {
						if stepType == 2 {
							date = start.AddDate(0, step*i, 0)
						} else {
							date = start.AddDate(0, 0, step*i)
						}
						if step >= 0 && date.Compare(stop) > 0 {
							break
						}
						if step < 0 && date.Compare(stop) < 0 {
							break
						}
						rst[i] = date.Format(time.DateOnly)
						i += 1
					}
				}
			default:
				err = errors.New("不支持的类型")
			}
			return
		},
	}
	tmpl.Funcs(funcMap)
	text := `
sequence: {{ sequence 1 10 1}}
sequence: {{ sequence 10 1 -2}}
sequence: {{ sequence "2023-04-17 20:00:00" "2023-05-07 10:00:00" "24h"}}
{{range $v := sequence "2023-04-27 20:00:00" "2023-05-07 10:00:00" "24h"}}
{{slice $v 0 10 -}}
{{end}}
sequence: {{ sequence "2023-04-27" "2023-05-07" "1d"}}
sequence: {{ sequence "2023-03-27" "2023-05-07" "1m"}}
sequence: {{ sequence "2023-03-27" "2023-05-27" "1m"}}
	`
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
