package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	palindrom "github.com/ermi9s/task2/go/Palindrom"
	frequency "github.com/ermi9s/task2/go/frequency"

)

func main() {
	for {

		reader := bufio.NewReader(os.Stdin)
		choice := 0
		HOME:
		fmt.Println("What Do you wnat:")
		fmt.Println("1.Word Frequncy Counter")
		fmt.Println("2.Palindrom Checker")
		fmt.Println("3. Exit")
		fmt.Print(">>>")
		input,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid Input!")
			goto HOME
		}
		num,err := strconv.Atoi(strings.TrimSpace(input));
		if err != nil {
			fmt.Println("Invalid Number Format!")
			goto HOME
		}
		choice = num;

		switch choice {
		case 1:
			frequency.Main()
		case 2:
			palindrom.Main();
		case 3:
			return
	}
}}