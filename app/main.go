package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/x/term"
	"github.com/codecrafters-io/shell-starter-go/app/evalutater"
	"github.com/codecrafters-io/shell-starter-go/app/parser"
	readline "github.com/codecrafters-io/shell-starter-go/app/readLine"
	"github.com/codecrafters-io/shell-starter-go/app/tokenizer"
)

func goRaw() {
	// Get current terminal state
	oldState, err := term.MakeRaw(os.Stdin.Fd())
	if err != nil {
		panic(err)
	}
	// Restore state on exit
	defer term.Restore(os.Stdin.Fd(), oldState)
}

func main() {
	goRaw()
	rd := readline.NewReadLine(os.Stdin, os.Stdout, os.Stderr)
	f, _ := os.Create("history.txt")
	f.Close()
	historyFile, err := os.OpenFile("history.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer historyFile.Close()
	for {
		fmt.Fprint(os.Stdout, "$ ")
		b, err := rd.Read()
		if err != nil {
			log.Fatal("error reading input")
		}
		historyFile.Write(b)
		b = b[:len(b)-1]
		if len(b) == 0 {
			continue
		}
		t := tokenizer.NewTokenizer(string(b))
		p := parser.NewParser(t)
		cmds := p.Parse()
		evalutater := evalutater.NewEvaluator()
		for _, stmt := range cmds.Statements {
			// fmt.Printf("Type of myInt: %T\n", stmt)
			evalutater.Eval(stmt, os.Stdin, os.Stdout, os.Stderr)
		}
	}
}
