package tokenizer

import (
	"github.com/codecrafters-io/shell-starter-go/app/token"
)

type Tokenizer struct {
	input   string
	pos     int
	word    string
	cmdRead bool
}

func NewTokenizer(input string) Tokenizer {
	t := Tokenizer{input: input}
	t.pos = 0
	t.skipWhilespace()
	t.readWord()
	return t
}

func (t *Tokenizer) readWord() {
	if t.pos >= len(t.input) {
		t.word = "\n"
		return
	}
	wrd := ""

	activeQt := ' '
	// read space separated words unless there are quotes then read ingore till quotes end
	for t.pos < len(t.input) && (t.input[t.pos] != ' ' || activeQt != ' ') {
		if t.input[t.pos] == '\'' || t.input[t.pos] == '"' {
			if activeQt != ' ' {
				if t.input[t.pos] == byte(activeQt) {
					activeQt = ' '
					t.pos += 1
					continue
				}
			} else {
				activeQt = rune(t.input[t.pos])
				t.pos += 1
				continue
			}
		}
		wrd += string(t.input[t.pos])
		t.pos += 1
	}
	t.word = wrd

}
func (t *Tokenizer) NextToken() token.Token {

	var tok token.Token
	t.skipWhilespace()
	if !t.cmdRead {
		switch t.word {
		case "exit":
			tok = newToken(token.EXIT, t.word)
		case "echo":
			tok = newToken(token.ECHO, t.word)
		case "type":
			tok = newToken(token.TYPE, t.word)
		case "pwd":
			tok = newToken(token.PWD, t.word)
		case "cd":
			tok = newToken(token.CD, t.word)
		case "cat":
			tok = newToken(token.CAT, t.word)
		default:
			tok = newToken(token.CUSTOM, t.word)
		}
		t.cmdRead = true
	} else {
		switch t.word {
		case "\n":
			tok = newToken(token.EOF, "")
		case ">":
			tok = newToken(token.GT, t.word)
		case "1>":
			tok = newToken(token.GT, t.word)
		case "2>":
			tok = newToken(token.GT2, t.word)
		case ">>":
			tok = newToken(token.RSHIFT, t.word)
		case "1>>":
			tok = newToken(token.RSHIFT, t.word)
		case "2>>":
			tok = newToken(token.RSHIFT2, t.word)
		case "|":
			tok = newToken(token.PIPE, t.word)
			t.cmdRead = false
		default:
			tok = newToken(token.ARG, t.word)
		}
	}
	t.readWord()
	return tok
}

func (t *Tokenizer) skipWhilespace() {
	for t.pos < len(t.input) && t.input[t.pos] == ' ' {
		t.pos += 1
	}
}

func newToken(typ token.TokenType, val string) token.Token {
	return token.Token{Type: typ, Val: val}
}
