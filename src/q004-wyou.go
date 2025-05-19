package main

import "fmt"


func can_change_str1_to_str2(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	// Create a map to count the frequency of characters in str1
	counts := make(map[rune]int)
	for _, char := range str1 {
		counts[char]++
	}
	// Decrease the frequency for characters in str2
	for _, char := range str2 {
		counts[char]--
		if counts[char] < 0 {
			// If any character in str2 has a higher frequency than in str1, return false
			return false
		}
	}
	// Check if all counts are zero
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	// If all counts are zero, str1 can be changed to str2
	return true
}

func main() {
	s1 := "This is golang"
	s2 := "gnalog si sihT"
	fmt.Println(can_change_str1_to_str2(s1, s2))

	s3 := "Here you are"
	s4 := "Are you here"
	fmt.Println(can_change_str1_to_str2(s3, s4))

	s5 := "This is golang1.1"
	s6 := "This is golang1"
	fmt.Println(can_change_str1_to_str2(s5, s6))
}
