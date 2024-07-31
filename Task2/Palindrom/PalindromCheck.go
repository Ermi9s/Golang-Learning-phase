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

		for l >= r {
			if input[l]!= input[r] {
				fmt.Println("Not Palindrom")
				return
			}
			l++
			r--
		}
		fmt.Println("The Word is Palindrom")
		return
	}

}