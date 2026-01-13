package handler

import (
	"bufio"
	"fmt"

	"github.com/number-system/internal/utils"
	"github.com/number-system/services"
)

func HandlerBaseConversion(scanner *bufio.Scanner) {
	for {
		fromBase, err := utils.InputBase(scanner, "Enter initial base: ")
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		toBase, err := utils.InputBase(scanner, "Enter destination base: ")
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		value, err := utils.InputValue(scanner, "Enter value: ", fromBase)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		result, err := services.ConvertFromBaseToBase(value, fromBase, toBase)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Result:", result)
		return
	}
}

func PauseAndClear(scanner *bufio.Scanner) {
	fmt.Print("Press 'Enter' to continue")
	scanner.Scan()
	utils.ClearScreen()
}
