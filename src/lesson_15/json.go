package main

import (
	"encoding/json"
	"fmt"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func main() {
	//sv := []string{"a", "b"}
	//boolVar, _ := json.Marshal(true)
	//svJ, _ := json.Marshal(sv)
	//fmt.Println(string(boolVar))
	//fmt.Println(string(svJ))
	//fmt.Println(svJ)

	//map1 := map[string]interface{}{
	//	"name":    "kira",
	//	"age":     33,
	//	"country": "japan",
	//}

	map1 := map[string]int{
		"name":    1,
		"age":     2,
		"country": 3,
	}

	map1Json, _ := json.Marshal(map1)
	fmt.Println(string(map1Json))

	data := []KeyValue{
		{Key: "name", Value: 1},
		{Key: "age", Value: 2},
		{Key: "country", Value: 3},
	}

	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))

	var data1 interface{}

	if err := json.Unmarshal(jsonData, &data1); err != nil {
		panic(err)
	}
	fmt.Println(data1)
}
