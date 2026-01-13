package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/number-system/handler"
	"github.com/number-system/utils"

	"github.com/number-system/services"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		services.PrintMenu()

		choice, err := utils.ReadProgram(scanner, "Select program: ", "Please enter correct program")
		if err != nil {
			fmt.Println("Please try again")
			continue
		}

		switch choice {
		case 1:
			handler.HandlerBaseConversion(scanner)
			handler.PauseAndClear(scanner)
		case 0:
			fmt.Println("Exit program")
			return
		default:
			fmt.Println("Feature not implemented yet")
			handler.PauseAndClear(scanner)
		}
	}
}

/*
In ordered
1. Enter fromBase
2. Enter toBase
3. Enter value (validate depend on "fromBase")
4. Convert
*/
