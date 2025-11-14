package token

type TokenType string

type Token struct {
	Type TokenType
	Val  string
}

const (
	ECHO       TokenType = "echo"
	EXIT       TokenType = "exit"
	TYPE       TokenType = "type"
	INVALIDCMD TokenType = "InvalidCmd"
	ARG        TokenType = "ARG"
	OPT        TokenType = "OPT"
	EOF        TokenType = "EOF"
)
