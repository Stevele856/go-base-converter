// Error handling
package utils

import (
	"fmt"
	"strconv"
)

/*Hàm validate binary
func ValidateBinary(input string) error {
	if input == "" {
		return fmt.Errorf("binary value is required")
	}
	regex := regexp.MustCompile(`^[01]+$`)

	if !regex.MatchString(input) {
		return fmt.Errorf("binary value must contain only 0 and 1")
	}

	return nil
}
*/

/*Hàm validate Decimal
func ValidateDecimal(input string) error {
	if input == ""{
		return fmt.Errorf("decimal value is required")
	}

	if !regexp.MustCompile(`^[0-9]+$`).MatchString(input) {
		return fmt.Errorf("decimal must contain number from 0-9")
	}

	return nil
}
*/

/* Hàm convert Character to Digit
xử lý 3 nhóm của HEX
| Input     | Kết quả |
| --------- | ------- |
| `'0'–'9'` | `0–9`   |
| `'A'–'F'` | `10–15` |
| `'a'–'f'` | `10–15` |
*/

func CharToDigit(ch rune) (int, error) {
	if ch >= '0' && ch <= '9' {
		return int(ch - '0'), nil
	}
	if ch >= 'A' && ch <= 'F' {
		return int(ch - 'A' + 10), nil
	}
	if ch >= 'a' && ch <= 'f' {
		return int(ch - 'a' + 10), nil
	}

	return 0, fmt.Errorf("invalid character value")
}

/*
Xử lý nhóm ngược lại
| Input     | Kết quả 	|
| --------- | ------- 	|
|   `0-9`  	| `'0'-'9'`	|
| 	`10-15`	| 	`A-F` 	|
*/

func DigitToChar(di int) (string, error) {
	if di < 0 || di > 15 {
		return "", fmt.Errorf("invalid Hex digit")
	}
	if di <= 9 {
		return strconv.Itoa(di), nil
	} else {
		return string('A' + rune(di-10)), nil
	}

}

// Hàm validate base
func ValidateBase(base int) error {
	if base != 2 && base != 8 && base != 10 && base != 16 {
		return fmt.Errorf("base value must be 2,8,10 or 16")
	}
	return nil
}

// Validate value theo base
func ValidateValue(value string, base int) error {
	if value == "" {
		return fmt.Errorf("value is required")
	}
	for _, ch := range value {
		digit, err := CharToDigit(ch)
		if err != nil {
			return fmt.Errorf("invalid character: %c", ch)
		}
		if digit >= base {
			return fmt.Errorf("digit %c is not valid for base %d", ch, base)
		}
	}
	return nil
}

/* 1011 -> 2*/
