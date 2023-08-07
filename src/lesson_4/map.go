package main

import "fmt"
import "github.com/mitchellh/mapstructure"

type Points struct {
	X int
	Y int
}

type Points1 struct {
	X int `mapstructure:"xx"`
	Y int `mapstructure:"yy"`
}

func main() {
	//RangeArr()
	//MapS()
	InToStruct()
}

func RangeArr() {
	arr := []int{1, 2, 3}
	for i := range arr {
		fmt.Println(i)
		fmt.Println(arr[i])
	}
}

func MapS() {
	pointsMap := map[string]Points{}
	intMap := make(map[int]Points)
	pointsMap["a"] = Points{1, 2}
	intMap[1] = Points{2, 3}
	fmt.Println(pointsMap)
	fmt.Println(intMap[1])
	fmt.Println(pointsMap["a"].X)

	value, ok := pointsMap["b"]
	fmt.Println(value)
	fmt.Println(ok)
}

func InToStruct() {
	map1 := map[string]int{
		"x": 1,
		"y": 2,
	}

	map2 := map[string]int{
		"xx": 3,
		"yy": 4,
	}

	p1 := Points{}
	p2 := Points1{}

	mapstructure.Decode(map1, &p1)
	fmt.Println(p1)

	mapstructure.Decode(map2, &p2)
	fmt.Println(p2)
}
