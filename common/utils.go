package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadIntInput(s string) int {
	for {
		fmt.Print(s)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		trimString := strings.TrimSpace(input)
		value, err := strconv.Atoi(trimString)
		if err != nil {
			fmt.Println("Kí tự không hợp lệ")
		}
		return value
	}
}

func ReadFloatInput(s string) float64 {
	for {
		fmt.Print(s)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		trimString := strings.TrimSpace(input)
		value, err := strconv.ParseFloat(trimString, 64)
		if err != nil {
			fmt.Println("Kí tự không hợp lệ")
		}
		return value
	}
}

func ReadStringInput(s string) string {
	fmt.Print(s)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	trimString := strings.TrimSpace(input)
	return trimString
}
