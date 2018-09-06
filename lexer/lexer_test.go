package lexer

import (
	"fmt"
	"testing"

	"../token"
)

type Expected struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	//input := `=+(){},;`
	tests := []Expected{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	for i, tt := range tests {
		fmt.Println(i, tt)
	}
}
