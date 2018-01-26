package util

import (
	"strings"
	"unicode/utf8"
)

func FormatLength(str string, length int) string {
	count := length - utf8.RuneCountInString(str)
	if count > 0 {
		str += strings.Repeat(" ", count)
	}
	return str
}
