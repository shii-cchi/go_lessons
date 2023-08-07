package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

func main() {
	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		} else if arg == "lower" {
			return strings.ToLower(json)
		}
		return json
	})

	json := `{"name": {"firstName": "сережа", "lastName": "сяонин"},"age": 24}`

	if !gjson.Valid(json) {
		panic("dont valid json")
	}

	value1 := gjson.Get(json, "name.lastName")
	value2 := gjson.Get(json, "name.firstName")
	fmt.Println("кто самый лучший малыш??? это", value2, value1)

	value3 := gjson.Get(json, `name.lastName|@case:upper`)
	fmt.Println(value3.String())

	fmt.Println(gjson.Parse(json).Get("age"))

	result, ok := gjson.Parse(json).Value().(map[string]interface{})

	if !ok {
		panic("parsing is not okey")
	}

	fmt.Println(result)
}
