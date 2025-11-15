package token

type TokenType string

type Token struct {
	Type TokenType
	Val  string
}

const (
	ECHO    TokenType = "echo"
	EXIT    TokenType = "exit"
	TYPE    TokenType = "type"
	PWD     TokenType = "pwd"
	CD      TokenType = "cd"
	CAT     TokenType = "cat"
	ARG     TokenType = "ARG"
	EOF     TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"
)
