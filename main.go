package main

import fmt "fmt"
import os "os"
import config "gitlab.kohn.io/ankoh/vmlcm/util"

func main() {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("You did not provide any arguments")
	} else {
		fmt.Println("Found the following arguments")
		fmt.Println(arguments)
	}
}
