package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a number between 1 and 10: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	num, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		fmt.Println("An error occured", err)
		return
	}

	fmt.Println("The number you entered is: ", num+1)
}
