package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	myString := "We-are back_again"

	newString := ToCamelCase(myString)

	fmt.Println(newString)
}

func ToCamelCase(s string) string {
	re := regexp.MustCompile(`[_\-\s]+`)
	phrase := re.Split(s, -1)

	firstWord := phrase[0]
	isUpperCase := isFirstLetterUpperCase(firstWord)

	for i, word := range phrase {

		if i == 0 && !isUpperCase {
			phrase[i] = strings.ToLower(word)
		} else {
			phrase[i] = capitalize(word)
		}
	}

	resultado := strings.Join(phrase, "")
	return resultado
}

func isFirstLetterUpperCase(word string) bool {
	if len(word) == 0 {
		return false
	}

	return unicode.IsUpper(rune(word[0]))
}

func capitalize(word string) string {
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])

	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}

	return string(runes)
}
