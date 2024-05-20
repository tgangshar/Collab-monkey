package parser 

import {

	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
}



type Parser struct {
	l *lexer.lexer // a pointer to an isntance of the lexer 

	curToken token.Token
	peekToken token.Token
}