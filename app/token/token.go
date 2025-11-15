package token

type TokenType string
type SubTokenType string

type Token struct {
	Type    TokenType
	SubType SubTokenType
	Val     string
}

const (
	CMD TokenType = "CMD"
	ARG TokenType = "ARG"
	OPT TokenType = "OPT"
	EOF TokenType = "EOF"
)

const (
	INVALIDCMD SubTokenType = "InvalidCmd"
	ECHO       SubTokenType = "echo"
	EXIT       SubTokenType = "exit"
	TYPE       SubTokenType = "type"
	PWD        SubTokenType = "pwd"
	CD         SubTokenType = "cd"
	CAT        SubTokenType = "cat"
)
