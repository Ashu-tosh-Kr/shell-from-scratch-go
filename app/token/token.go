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
	HISTORY TokenType = "history"
	CUSTOM  TokenType = "CUSTOM"
	ARG     TokenType = "ARG"
	GT      TokenType = ">"
	GT2     TokenType = "2>"
	RSHIFT  TokenType = ">>"
	RSHIFT2 TokenType = "2>>"
	LT      TokenType = "<"
	PIPE    TokenType = "|"
	EOF     TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"
)

func IsCmd(typ TokenType) bool {
	switch typ {
	case ECHO, TYPE, PWD, EXIT, CD, CAT, CUSTOM:
		return true
	default:
		return false

	}
}
