package util

import (
	"unicode/utf8"
	"strings"
)

func FormatLength(str string, length int) string {
	count := length - utf8.RuneCountInString(str)
	if count>0 {
		str += strings.Repeat(" ",count)
	}
	return str
}