package rhombifer

import (
	"testing"

	"github.com/racg0092/rhombifer/lexer"
	"github.com/racg0092/rhombifer/tokens"
)

func TestNextToken(t *testing.T) {
	input := `proxy set --name cool --lang go -p --url go.codeh.io`

	tests := []tokens.Token{
		{tokens.IDENT, "proxy"},
		{tokens.IDENT, "set"},
		{tokens.DOUBLE_DASH, "--"},
		{tokens.IDENT, "name"},
		{tokens.IDENT, "cool"},
		{tokens.DOUBLE_DASH, "--"},
		{tokens.IDENT, "lang"},
		{tokens.IDENT, "go"},
		{tokens.DASH, "-"},
		{tokens.IDENT, "p"},
		{tokens.DOUBLE_DASH, "--"},
		{tokens.IDENT, "url"},
		{tokens.IDENT, "go.codeh.io"},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.Type {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", i, tt.Type, tok.Type)
		}

		if tok.Literal != tt.Literal {
			t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q", i, tt.Literal, tok.Literal)
		}
	}
}
