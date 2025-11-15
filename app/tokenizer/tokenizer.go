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

	// read quoted words
	if t.pos < len(t.input) && (t.input[t.pos] == '"' || t.input[t.pos] == '\'') {
		delim := t.input[t.pos]
		t.pos += 1
		for t.pos < len(t.input) && t.input[t.pos] != delim {
			wrd += string(t.input[t.pos])
			t.pos += 1
		}
		t.pos += 2
		t.word = wrd
		return
	}

	// read space separated words
	for t.pos < len(t.input) && t.input[t.pos] != ' ' {

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
			tok = newToken(token.CMD, token.EXIT, t.word)
		case "echo":
			tok = newToken(token.CMD, token.ECHO, t.word)
		case "type":
			tok = newToken(token.CMD, token.TYPE, t.word)
		case "pwd":
			tok = newToken(token.CMD, token.PWD, t.word)
		case "cd":
			tok = newToken(token.CMD, token.CD, t.word)
		case "cat":
			tok = newToken(token.CMD, token.CAT, t.word)
		default:
			tok = newToken(token.CMD, token.INVALIDCMD, t.word)
		}
		t.cmdRead = true
	} else {
		if t.word[0] == '-' {
			tok = newToken(token.OPT, "", t.word)
		} else if t.word == "\n" {
			tok = newToken(token.EOF, "", "")
		} else {
			tok = newToken(token.ARG, "", t.word)
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

func newToken(typ token.TokenType, subTyp token.SubTokenType, val string) token.Token {
	return token.Token{Type: typ, SubType: subTyp, Val: val}
}
