package parser

import (
	"github.com/codecrafters-io/shell-starter-go/app/ast"
	"github.com/codecrafters-io/shell-starter-go/app/token"
	"github.com/codecrafters-io/shell-starter-go/app/tokenizer"
)

type Parser struct {
	t         tokenizer.Tokenizer
	curToken  token.Token
	peekToken token.Token
}

func NewParser(t tokenizer.Tokenizer) *Parser {
	p := &Parser{
		t: t,
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Parse() ast.Commands {
	ast := ast.Commands{Statements: []ast.BaseCmd{}}
	for p.curToken.Type != token.EOF {

		stmt := p.parseStatement()
		ast.Statements = append(ast.Statements, stmt)
	}
	return ast
}

func (p *Parser) parseStatement() ast.BaseCmd {
	var stmt any
	for p.curToken.Type != token.EOF {
		if token.IsCmd(p.curToken.Type) {
			stmt = p.parseCmd()
		}
		if p.curToken.Type == token.GT {
			stmt = p.parseRedirectCmd(stmt)
		}
		if p.curToken.Type == token.PIPE {
			stmt = p.parsePipedCmd(stmt)
		}
	}
	return stmt
}

func (p *Parser) parseCmd() ast.BaseCmd {
	cmd := ast.SimpleCmd{}
	cmd.Cmd = p.curToken
	p.nextToken()
	for p.curToken.Type != token.EOF && p.curToken.Type == token.ARG {
		cmd.Args = append(cmd.Args, p.curToken)
		p.nextToken()
	}
	return cmd
}

func (p *Parser) parseRedirectCmd(cmd ast.BaseCmd) ast.RedirectCmd {
	redir := ast.RedirectCmd{Cmd: cmd}
	if p.curToken.Type == token.GT {
		redir.RedirStdOut = true
	}
	if p.curToken.Type == token.GT2 {
		redir.RedirStdErr = true
	}
	p.nextToken()
	redir.RedirectTo = p.curToken
	p.nextToken()
	return redir
}

func (p *Parser) parsePipedCmd(cmd ast.BaseCmd) ast.PipedCmd {
	piped := ast.PipedCmd{Left: cmd}
	p.nextToken()
	piped.Right = p.parseCmd()
	return piped
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.t.NextToken()
}
