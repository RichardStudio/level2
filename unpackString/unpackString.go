package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(s string) (string, error) {
	var result strings.Builder
	var escape bool

	for i, r := range s {
		if unicode.IsDigit(r) && !escape {
			if i == 0 || unicode.IsDigit(rune(s[i-1])) {
				return "", errors.New("invalid string")
			}
			count, err := strconv.Atoi(string(r))
			if err != nil {
				return "", err
			}
			result.WriteString(strings.Repeat(string(s[i-1]), count-1))
		} else if r == '\\' && !escape {
			escape = true
		} else {
			result.WriteRune(r)
			escape = false
		}
	}
	return result.String(), nil
}

func main() {
	tests := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		`qwe\4\5`,
		`qwe\\5`,
	}

	for _, test := range tests {
		result, err := Unpack(test)
		if err != nil {
			fmt.Printf("Error unpacking %q: %v\n", test, err)
		} else {
			fmt.Printf("Unpacked %q: %q\n", test, result)
		}
	}
}
