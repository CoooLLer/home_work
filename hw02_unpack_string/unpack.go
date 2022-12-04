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
	var previousChar rune

	for pos, char := range str {
		if pos == 0 {
			if unicode.IsDigit(char) {
				return "", ErrInvalidString
			}
			previousChar = char
			continue
		}

		if unicode.IsDigit(char) {
			if unicode.IsDigit(previousChar) {
				return "", ErrInvalidString
			}
			repeatCount, err := strconv.Atoi(string(char))
			if err != nil {
				return "", ErrInvalidString
			}

			builder.WriteString(strings.Repeat(string(previousChar), repeatCount))
			previousChar = char
			continue
		}

		if !unicode.IsDigit(previousChar) {
			builder.WriteString(string(previousChar))
		}

		previousChar = char

		if pos+1 == len(str) {
			builder.WriteString(string(char))
		}
	}

	return builder.String(), nil
}
