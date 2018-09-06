package lexer

import (
	"testing"
	"Compiler/token"
)

type Expected struct {
	expectedType token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T){
	input := `=+(){},;`
	tests := []Expected{
		{token.ASSIGN,"="}
	}
}
