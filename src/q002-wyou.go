package main

import "fmt"
import "strings"

func contain_duplicated_chars(str string) bool {
	// Create a map to track the characters we've seen
	charMap := make(map[rune]bool)

	// Iterate over each character in the string
	for _, char := range str {
		// If the character is already in the map, return true (duplicate found)
		if charMap[char] {
			return true
		}
		// Otherwise, add the character to the map
		charMap[char] = true
	}

	// If we finish iterating without finding duplicates, return false
	return false
}

func contain_duplicated_chars2(str string) bool {
	if (len(str) > 3000) {
		fmt.Println("The string is longer than 3000.")
		return false
	}

	for _, c := range str {
		if strings.Count(str, string(c)) > 1 {
			return true
		}
	}

	return false
}

func main() {
	var str_to_test = "asdfg";
	if contain_duplicated_chars2(str_to_test) {
		fmt.Println("The string contains duplicated characters.")
	} else {
		fmt.Println("The string does not contain duplicated characters.")
	}
}