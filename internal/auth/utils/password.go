package utils

import (
	"strings"
	"unicode/utf8"
)

func CheckPasswordsEquality(pwd1, pwd2 string) bool {
	return strings.EqualFold(pwd1, pwd2)
}

func CheckPasswordSecurity(pwd string) bool {
	return (utf8.RuneCountInString(pwd) > 8)
}
