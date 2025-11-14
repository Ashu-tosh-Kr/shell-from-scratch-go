package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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
				fmt.Printf("%s is a shell builtin", val.Val)
			default:
				p := os.Getenv("PATH")
				dirs := strings.Split(p, string(os.PathListSeparator))
				found := false
				for _, dir := range dirs {
					info, err := os.Stat(fmt.Sprintf("%s/%s", dir, val.Val))
					if err != nil {
						continue
					}
					if info.IsDir() {
						continue
					}
					md := info.Mode()
					isExecutable := md&0111 != 0
					if !isExecutable {
						continue
					}
					fmt.Printf("%s is %s/%s", val.Val, dir, val.Val)
					found = true
					break
				}
				if !found {
					fmt.Printf("%s: not found", val.Val)
				}

			}
			fmt.Println()

		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", mainCmd.Val)
		}
	}
}
