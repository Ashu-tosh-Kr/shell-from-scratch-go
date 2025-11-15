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
	CUSTOM  TokenType = "CUSTOM"
	ARG     TokenType = "ARG"
	GT      TokenType = ">"
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
