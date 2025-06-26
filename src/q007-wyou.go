package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

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

type People struct {
	Name string `json:"name"`
}

func (p *People) String() string {
	//return fmt.Sprintf("print: %v", p) // %v will invoke the String method of (*People)
	return fmt.Sprintf("print: %v", *p) // *p convert *People to People, then the interface func (p *People) String() will NOT be invoked
	//return fmt.Sprintf("String method: Name=%s", p.Name)
}

func wyou_sleep() {
	ch := make(chan int, 1000)
	var wg sync.WaitGroup

	// Start sender
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// Close channel after sending all values
		close(ch)
	}()

	// Start receiver
	go func() {
		for a := range ch {
			fmt.Println("a: ", a)
		}
		fmt.Println("channel closed")
	}()

	// Wait for sender to finish and close channel
	wg.Wait()
	fmt.Println("ok")
	time.Sleep(time.Second * 1)
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

	var p People
	js := `{
		"name":"11"
	}`
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Printf("people: %s", p.Name)
	fmt.Printf("people: %v", p)
	fmt.Println(p.String())

	wyou_sleep()
}
