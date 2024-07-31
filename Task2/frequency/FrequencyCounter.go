package frequency

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
	mapping := map[string]int{"":0}

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

		inputbits := strings.Split(input , " ")

		for _,val := range inputbits{
			mapping[strings.TrimSpace((val))]++;
		}

		fmt.Println("Final Count")
		for key,val := range mapping{
			if key != " "{
				fmt.Println("Word: ", key , "Count: ",val)
			}	
			
		}
		break
	}

}