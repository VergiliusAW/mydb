package main

import (
	"fmt"
	"mydb/helpers"
	"os"
)

func main() {
	argsWithProg := os.Args
	if len(argsWithProg) < 3 {
		fmt.Println(errArgs())
		return
	}
	args := os.Args[1:]

	switch args[0] {
	case "-c":
		if len(args) == 2 {
			helpers.CopyToClipboard(args[1])
		} else {
			fmt.Println(errArgs())
		}
	case "--get", "-G":
		if len(args) == 2 {
			fmt.Println(helpers.Get(args[1]))
		} else {
			fmt.Println(errArgs())
		}
	case "--set", "-S":
		if len(args) == 3 {
			helpers.Set(args[1], args[2])
		} else {
			fmt.Println(errArgs())
		}
	default:
		fmt.Println(errArgs())
	}

}

func errArgs() string {
	return "usage: mydb [--get <name>] [-G <name>] [-set <name> <value>] [-S <name> <value>]"
}
