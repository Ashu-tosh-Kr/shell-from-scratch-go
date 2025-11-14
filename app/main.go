package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/codecrafters-io/shell-starter-go/app/token"
	"github.com/codecrafters-io/shell-starter-go/app/tokenizer"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		b, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
		if err != nil {
			log.Fatal("error reading input")
		}
		b = b[:len(b)-1]
		t := tokenizer.NewTokenizer(string(b))
		mainCmd := t.NextToken()
		switch mainCmd.SubType {
		case token.EXIT:
			arg := t.NextToken()
			if arg.Type != token.ARG {
				fmt.Fprint(os.Stdout, "invalid code")
			}
			v, err := strconv.Atoi(arg.Val)
			if err != nil {
				fmt.Fprint(os.Stdout, "invalid code")
			}
			os.Exit(v)

		case token.ECHO:
			for {
				val := t.NextToken()
				if val.Type == token.EOF {
					break
				}
				fmt.Print(val.Val + " ")
			}
			fmt.Println()
		case token.TYPE:
			val := t.NextToken()
			typ := token.SubTokenType(val.Val)
			switch typ {
			case token.ECHO, token.EXIT, token.TYPE:
				fmt.Printf("%s is a shell builin", val.Val)
			default:
				fmt.Printf("%s: not found", val.Val)

			}
			fmt.Println()

		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", mainCmd.Val)
		}
	}
}
