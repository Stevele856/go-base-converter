package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ReadProgram(scanner *bufio.Scanner, prompt, errMsg string) (int, error) {
	regex := regexp.MustCompile(`[0-1]+$`)

	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			fmt.Println(errMsg)
			continue
		}

		if !regex.MatchString(input) {
			fmt.Println("Incorrect program")
			continue
		}

		return strconv.Atoi(input)
	}
}