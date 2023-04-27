package json

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

type Student struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address *Address `json:"address"`
	Scores  []*Score `json:"scores"`
}

type Address struct {
	Province string `json:"province"`
	City     string `json:"city"`
}

type Score struct {
	Course string  `json:"course"`
	Score  float32 `json:"score"`
}

func TestMarshalStruct(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	student := Student{
		Name:    "小明",
		Age:     20,
		Address: &Address{Province: "河北", City: "邯郸"},
		Scores:  []*Score{{Course: "english", Score: 80}, {Course: "math", Score: 90}},
	}
	bytes, _ := json.Marshal(student)
	fmt.Println(string(bytes))

	str, _ := json.MarshalToString(student)
	fmt.Println(str)

	str, _ = json.MarshalToString(&student)
	fmt.Println(str)
}

func TestMarshalMap(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	student := map[string]interface{}{
		"name":    "小明",
		"age":     20,
		"address": map[string]interface{}{"province": "河北", "city": "邯郸"},
	}

	bytes, _ := json.Marshal(student)
	fmt.Println(string(bytes))

	str, _ := json.MarshalToString(student)
	fmt.Println(str)

	str, _ = json.MarshalToString(&student)
	fmt.Println(str)
}

func TestUnmarshalStruct(t *testing.T) {
	str := `{"name":"小明","age":20,"address":{"province":"河北","city":"邯郸"},"scores":[{"course":"english","score":80},{"course":"math","score":90}]}`
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	student := Student{}
	json.UnmarshalFromString(str, &student)
	fmt.Println(student)

	student = Student{}
	json.Unmarshal([]byte(str), &student)
	fmt.Println(student)
}

func TestUnmarshalMap(t *testing.T) {
	str := `{"name":"小明","age":20,"address":{"province":"河北","city":"邯郸"},"scores":[{"course":"english","score":80},{"course":"math","score":90}]}`
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	student := map[string]interface{}{}
	json.UnmarshalFromString(str, &student)
	fmt.Println(student)

	student = map[string]interface{}{}
	json.Unmarshal([]byte(str), &student)
	fmt.Println(student)
}
