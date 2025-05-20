package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	s := Show{
		Param: map[string]interface{}{
			"key": "value",
		},
	}

	s2 := Show{
		Param: make(map[string]interface{}),
	}
	s2.Param["key2"] = "value2"

	fmt.Println(s.Param["key"])
	fmt.Println(s2.Param["key2"])
}