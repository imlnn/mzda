package utils

import (
	"strings"
	"unicode/utf8"
)

const PWD_MIN_LEN = 8

func CheckPasswordsEquality(pwd1, pwd2 string) bool {
	const fn = "internal/utils/password/CheckPasswordsEquality"
	return strings.EqualFold(pwd1, pwd2)
}

func CheckPasswordSecurity(pwd string) bool {
	const fn = "internal/utils/password/CheckPasswordSecurity"
	return utf8.RuneCountInString(pwd) > PWD_MIN_LEN
}
