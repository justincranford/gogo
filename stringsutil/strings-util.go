package stringsutil

import (
	"fmt"
	"strings"
	"time"
	"unsafe"
)

var PalindromeTestStrings = []string{
	"racecar",
	"Hello",
	"madam",
	"A man a plan a canal Panama",
	"not a palindrome",
}

func ReverseString(s string) string {
	reverse := make([]rune, len(s))
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		reverse[left] = rune(s[right])
		reverse[right] = rune(s[left])
	}
	for index, arune := range s {
		reverse[len(s)-1-index] = arune
	}
	return string(reverse)
}

func ReverseStringsSlice(str []string) []string {
	reversed := make([]string, len(str))
	for left, right := 0, len(str)-1; left <= right; left, right = left+1, right-1 {
		reversed[left], reversed[right] = str[right], str[left]
	}
	return reversed
}

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return false
		}
	}
	return true
}

func PrintTests() (*[5]int, uintptr) {

	var a = [5]int{1, 3, 5, 7, 9}
	var b, c = len(a), "fish"
	d := &a
	pointerSize := unsafe.Sizeof(d)
	fmt.Println(a, b, c, d)

	fmt.Printf("Pointer address: %016x\n", d)
	fmt.Printf("Pointer pointerSize: %d\n", pointerSize)

	return d, pointerSize
}

func SwitchTests() {
	for i := 0; i < 4; i++ {
		switch i {
		case 1, 3:
			fmt.Println("odd")
		case 0, 2:
			fmt.Println("even")
		}
		switch i {
		case 1:
			fmt.Println("one")
		case 0:
			fmt.Println("zero")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		}
		fmt.Printf("i: %d\n", i)
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Yah! O")
	default:
		fmt.Println("Boo!")
	}
}

func PalindromeTests() {
	originalString := "Hello, World!"
	reversedString := ReverseString(originalString)
	fmt.Println("originalString:", originalString)
	fmt.Println("reversedString:", reversedString)

	inputString := "This is a test string"
	originalTokens := strings.Split(inputString, " ")
	reversedTokens := ReverseStringsSlice(originalTokens)
	fmt.Println("originalTokens:", originalTokens)
	fmt.Println("reversedTokens:", reversedTokens)

	for _, testString := range PalindromeTestStrings {
		yes := IsPalindrome(testString)
		fmt.Println("testString:", testString, yes)
	}
}
