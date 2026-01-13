package services

import (
	"fmt"

	"github.com/number-system/utils"
)

func BaseToDecimal(value string, base int) (int, error) {
	if value == "" {
		return 0, fmt.Errorf("value is required")
	}
	result := 0

	for _, r := range value {
		digit, err := utils.CharToDigit(r)
		if err != nil {
			return 0, err
		}
		if digit >= base {
			return 0, fmt.Errorf("digit %c is not valid for base %d", r, base)
		}
		result = result*base + digit
	}
	return result, nil
}


func DecimalToBase(value, base int) (string, error) {
	if base < 2 || base > 16 {
		return "", fmt.Errorf("base must be between 2 and 16")
	}

	if value == 0 {
		return "0", nil
	}

	result := ""

	for value > 0 {
		remain := value % base

		str, err := utils.DigitToChar(remain)
		if err != nil {
			return "", err
		}

		result = str + result
		value /= base
	}

	return result, nil
}


func ConvertFromBaseToBase(value string, fromBase, toBase int) (string, error) {
	decimal, err := BaseToDecimal(value, fromBase)

	if err != nil {
		return "", err
	}
	return DecimalToBase(decimal, toBase)
}
