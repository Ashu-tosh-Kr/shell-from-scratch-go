package main

import (
	"log"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/evalutater"
	"github.com/codecrafters-io/shell-starter-go/app/parser"
	readline "github.com/codecrafters-io/shell-starter-go/app/readLine"
	"github.com/codecrafters-io/shell-starter-go/app/tokenizer"
)

func main() {
	rd := readline.NewReadLine(os.Stdin, os.Stdout, os.Stderr)

	for {
		b, err := rd.Read()
		if err != nil {
			log.Fatal("error reading input")
		}
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
