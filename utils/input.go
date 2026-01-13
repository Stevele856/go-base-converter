package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Check Input value
func InputValue(scanner *bufio.Scanner, prompt string, base int) (string, error) {
	for {
		fmt.Print(prompt)

		if !scanner.Scan() {
			return "", fmt.Errorf("failed to read input")
		}

		input := strings.TrimSpace(scanner.Text())

		if err := ValidateValue(input, base); err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		return input, nil
	}
}

// Check input base
func InputBase(scanner *bufio.Scanner, prompt string) (int, error) {
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			return 0, fmt.Errorf("cannot read base value")
		}
		input := strings.TrimSpace(scanner.Text())
		if input == ""{
			fmt.Println("Base value is required")
			continue
		}

		base, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("base must be a number")
			continue
		}

		if err := ValidateBase(base); err != nil {
			fmt.Println(err.Error())
			continue
		}
		return base, nil
	}
}
