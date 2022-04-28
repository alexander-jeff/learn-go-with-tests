package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alexander-jeff/learn-go-with-tests/hello"
)

func main() {
	args := os.Args[1:]
	size := len(args)

	switch size {
		case 0:
			fmt.Println("NO ARGS at least one arg required")
		case 1:
			name := args[0]
			fmt.Println(hello.Hello(name, ""))
		case 2:
			name := args[0]
			lang := args[1]
			fmt.Println(hello.Hello(name, lang))
		default:
			progname := os.Args[0]
			got := strings.Join(os.Args, " ")
			want := progname + " name [language]"
			fmt.Println("Too many arguments.")
			fmt.Println("Got " + got)
			fmt.Println("Want " + want)
	}
}
