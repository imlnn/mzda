package utils

import (
	"fmt"
	"strings"
	"unicode"
)

const PWD_MIN_LEN = 8

func CheckPasswordsEquality(pwd1, pwd2 string) bool {
	const fn = "internal/utils/password/CheckPasswordsEquality"
	return strings.EqualFold(pwd1, pwd2)
}

func CheckPasswordSecurity(pwd string) error {
	const fn = "internal/utils/password/CheckPasswordSecurity"

	hasLowerCase := false
	hasUpperCase := false
	hasDigit := false

	if len(pwd) >= PWD_MIN_LEN {
		for _, char := range pwd {
			if unicode.IsLower(char) {
				hasLowerCase = true
			}
			if unicode.IsUpper(char) {
				hasUpperCase = true
			}
			if unicode.IsDigit(char) {
				hasDigit = true
			}
			if hasDigit && hasLowerCase && hasUpperCase {
				break
			}
		}
	}

	if len(pwd) < PWD_MIN_LEN {
		return fmt.Errorf("password have len < 8")
	}

	if !hasLowerCase {
		return fmt.Errorf("password doesn't contain lower case chars")
	}

	if !hasUpperCase {
		return fmt.Errorf("password doesn't contain upper case chars")
	}

	if !hasDigit {
		return fmt.Errorf("password doesn't contain digits")
	}

	return nil
}
