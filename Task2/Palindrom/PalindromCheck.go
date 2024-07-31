package Palindrom

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func sanitize(word string) string {
	var final string
	for _,val := range word {
		if (int32(val) - int32('a') >= 0 && int32(val) - int32('a') <= 26) || val == ' ' {
			final += string(val)
		}
	}

	return final
}

func Main() {

	for {
		fmt.Print("INPUT A STRING: ")
		reader := bufio.NewReader(os.Stdin)
		input,err := reader.ReadString('\n')
		
		if err != nil {
			fmt.Println("Invalid Input!")
			continue
		}

		input = strings.ToLower(input)
		input = sanitize(input)
		input = strings.TrimSpace(input)

		l:= 0
		r:= len(input)-1

		const (
			Reset  = "\033[0m"
			Yellow = "\033[33m"
			Red    = "\033[31m"
		)
		for l <= r {
			if input[l] != input[r] {
				fmt.Println(Red,"Not Palindrom",Reset)
				return
			}
			l++
			r--
		}
		fmt.Println(Yellow , "The Word is Palindrom" , Reset)
		return
	}

}