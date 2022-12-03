package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var builder strings.Builder
	skipNextChar := false

	for pos, char := range str {
		if skipNextChar {
			skipNextChar = false
			continue
		}

		if unicode.IsDigit(char) {
			return "", ErrInvalidString
		}

		if pos+1 < len(str) {
			nextChar := rune(str[pos+1])
			if unicode.IsDigit(nextChar) {
				i, err := strconv.Atoi(string(nextChar))
				if err != nil {
					panic(err)
				}
				builder.WriteString(strings.Repeat(string(char), i))

				skipNextChar = true

				continue
			}
		}

		builder.WriteString(string(char))
	}

	return builder.String(), nil
}
