package main

import "fmt"

func string_reverse(s string) string {
	str := []rune(s)
	l := len(str)

	for i := 0; i < l/2; i++ {
		str[i], str[l-i-1] = str[l-i-1], str[i]
	}

	return string(str);
}

func main() {
	str_to_be_reversed := "This is golang"
	reversed_str := string_reverse(str_to_be_reversed)
	fmt.Println(str_to_be_reversed)
	fmt.Println(reversed_str)
}