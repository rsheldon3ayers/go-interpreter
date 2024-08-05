package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/rsheldon3ayers/go-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	firstName := strings.Split(user.Name, " ")[0]
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		firstName)
	fmt.Printf("Feel free to type commands \n")
	repl.Start(os.Stdin, os.Stdout)
}
