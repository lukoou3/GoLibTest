package json

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"testing"
)

func TestJsonPath(t *testing.T) {
	str := `{"name":"小明","age":20,"address":{"province":"河北","city":"邯郸"},"scores":[{"course":"english","score":80},{"course":"math","score":90}]}`
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	student := map[string]interface{}{}
	json.UnmarshalFromString(str, &student)
	fmt.Println(student)

	obj, _ := oj.ParseString(str)
	x, _ := jp.ParseString("$.name")
	name := x.Get(obj)
	fmt.Println(name)

	x, _ = jp.ParseString("$.age")
	age := x.Get(obj)
	fmt.Println(age)

	x, _ = jp.ParseString("$.address.province")
	province := x.Get(obj)
	fmt.Println(province)
	fmt.Println(x.Get(&obj)) // 传入指针得不到结果

	x, _ = jp.ParseString("$.scores[?(@.course == 'math')]")
	math := x.Get(obj)
	fmt.Println(math)
	fmt.Println(x.Get(&obj)) // 传入指针得不到结果

	fmt.Printf("%T\n", obj)  // map[string]interface {}
	fmt.Printf("%T\n", &obj) // *interface {}
}

func TestJsonPathJsoniter(t *testing.T) {
	str := `{"name":"小明","age":20,"address":{"province":"河北","city":"邯郸"},"scores":[{"course":"english","score":80},{"course":"math","score":90}]}`
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	student := map[string]interface{}{}
	json.UnmarshalFromString(str, &student)
	fmt.Println(student)

	obj := map[string]interface{}{}
	json.UnmarshalFromString(str, &obj)
	x, _ := jp.ParseString("$.name")
	name := x.Get(obj)
	fmt.Println(name)

	x, _ = jp.ParseString("$.age")
	age := x.Get(obj)
	fmt.Println(age)

	x, _ = jp.ParseString("$.address.province")
	province := x.Get(obj)
	fmt.Println(province)
	fmt.Println(x.Get(&obj)) // 可以得到结果

	x, _ = jp.ParseString("$.scores[?(@.course == 'math')]")
	math := x.Get(obj)
	fmt.Println(math)
	fmt.Println(x.Get(&obj)) // 可以得到结果

	fmt.Printf("%T\n", obj)  // map[string]interface {}
	fmt.Printf("%T\n", &obj) // *map[string]interface {}
}
