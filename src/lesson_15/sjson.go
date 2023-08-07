package main

import (
	"fmt"
	"github.com/tidwall/sjson"
)

func main() {
	json := `{"name": {"firstName": "сережа", "lastName": "сяонин"},"age": 24, "hobbies": ["обижать чопика", "игнорить чопика", "бить чопика по жопе"]}`

	json, _ = sjson.Set(json, "name.firstName", "kisik")
	fmt.Println(json)

	json, _ = sjson.Set(json, "hobbies.-1", "любить чопика")
	fmt.Println(json)

	json1, _ := sjson.Set("", "name", "chopik")
	fmt.Println(json1)

	val, _ := sjson.Delete(json, "age")
	fmt.Println(val)
}
