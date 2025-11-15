package ast

import "github.com/codecrafters-io/shell-starter-go/app/token"

type BaseCmd any

type SimpleCmd struct {
	Cmd  token.Token
	Args []token.Token
}

type RedirectCmd struct {
	Cmd         BaseCmd
	RedirectTo  token.Token
	RedirStdErr bool
	RedirStdOut bool
	AppendMode  bool
}

type PipedCmd struct {
	Left  BaseCmd
	Right BaseCmd
}

type Commands struct {
	Statements []BaseCmd
}
