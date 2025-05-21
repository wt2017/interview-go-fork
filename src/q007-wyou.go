package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	Param
}

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case student:
		fmt.Println("struct type", msg.Name)
	case *student:
		fmt.Println("point type", msg.Name)
	}
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

	zhoujielun(student{Name: "zhoujielun"})
	zhoujielun(&student{Name: "zhoujielun"})
}