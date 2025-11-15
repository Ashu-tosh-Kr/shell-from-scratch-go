package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/ast"
	"github.com/codecrafters-io/shell-starter-go/app/parser"
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
		p := parser.NewParser(t)
		cmds := p.Parse()
		for _, stmt := range cmds.Statements {
			// fmt.Printf("Type of myInt: %T\n", stmt)
			fmt.Print(eval(stmt))
		}
	}
}

func eval(stmt ast.BaseCmd) string {
	switch stmt := stmt.(type) {

	case ast.SimpleCmd:
		switch stmt.Cmd.Type {
		case token.EXIT:
			if len(stmt.Args) > 1 {
				return "exit: too many arguments"
			}
			v, err := strconv.Atoi(stmt.Args[0].Val)
			if err != nil {
				fmt.Fprint(os.Stdout, "invalid code")
			}
			os.Exit(v)

		case token.ECHO:
			var output string
			for _, arg := range stmt.Args {

				if arg.Type == token.EOF {
					break
				}
				output += arg.Val + " "
			}
			output += "\n"
			return output
		case token.TYPE:
			var output string
			for _, arg := range stmt.Args {

				typ := token.TokenType(arg.Val)
				switch typ {
				case token.ECHO, token.EXIT, token.TYPE, token.PWD:
					output += fmt.Sprintf("%s is a shell builtin\n", arg.Val)
				default:
					path, found := findProgInPath(arg.Val)
					if !found {
						output += fmt.Sprintf("%s: not found\n", arg.Val)
					} else {
						output += fmt.Sprintf("%s is %s/%s\n", arg.Val, path, arg.Val)
					}

				}
			}
			return output

		case token.PWD:
			if len(stmt.Args) != 0 {
				return fmt.Sprintln("pwd: too many arguments")
			}
			path, _ := os.Getwd()
			return fmt.Sprintln(path)

		case token.CD:
			if len(stmt.Args) == 0 {
				stmt.Args = append(stmt.Args, token.Token{Type: token.ARG, Val: "~"})
			}
			homedir, _ := os.UserHomeDir()
			stmt.Args[0].Val = strings.Replace(stmt.Args[0].Val, "~", homedir, 1)
			err := os.Chdir(stmt.Args[0].Val)
			if err != nil {
				return fmt.Sprintf("cd: %s: No such file or directory\n", stmt.Args[0].Val)
			}
			return ""
		case token.CAT:
			var finOut string
			for _, arg := range stmt.Args {
				cmd := exec.Command("cat", arg.Val)
				output, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Print(err.Error())
				}
				finOut += string(output)
			}
			return finOut

		default:
			_, ok := findProgInPath(stmt.Cmd.Val)
			if !ok {
				return fmt.Sprintf("%s: command not found\n", stmt.Cmd.Val)
			}
			optAndArgs := make([]string, 0)
			for _, arg := range stmt.Args {
				optAndArgs = append(optAndArgs, arg.Val)
			}
			cmd := exec.Command(stmt.Cmd.Val, optAndArgs...)
			output, _ := cmd.CombinedOutput()

			return fmt.Sprint(string(output))
		}
	case ast.PipedCmd:
		_ = eval(stmt.Left)
		return eval(stmt.Right)

	}
	return "Something went wrong"
}

func findProgInPath(prog string) (string, bool) {
	var path string
	p := os.Getenv("PATH")
	dirs := strings.Split(p, string(os.PathListSeparator))
	found := false
	for _, dir := range dirs {
		info, err := os.Stat(fmt.Sprintf("%s/%s", dir, prog))
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
		path = dir
		found = true
		break
	}
	return path, found
}
