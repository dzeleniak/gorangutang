package main

import (
	"fmt"
	"gorangutang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s...\n\nWelcome to Gorangutang\n", user.Username)

	repl.StartLexer(os.Stdin, os.Stdout)

}
