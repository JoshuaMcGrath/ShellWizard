package main

import (
	"fmt"
	"shell_wizard/input_handler"
)

func main() {
	//fmt.Println("The beginning of your journey")
	//command, args := input_handler.RequestInput()
	//output, err := input_handler.RunInput(command, args)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(output)
	count := 89
	for count != 100 {
		//command := input_handler.RequestInput()
		//output := input_handler.PipeResolver(command)
		output, _ := input_handler.TryBash()
		fmt.Println(output)
	}
}
