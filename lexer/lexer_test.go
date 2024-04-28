package lexer

import (
	"gorangutang/token"
	"testing"
)

type LexerTestCase struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func (tc LexerTestCase) Check(l *Lexer, i int, t *testing.T) {
	tok := l.NextToken()

	if tok.Type != tc.expectedType {
		t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
			i, tc.expectedType, tok.Type)
	}

	if tok.Literal != tc.expectedLiteral {
		t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
			i, tc.expectedLiteral, tok.Literal)
	}
}

func TestLexer_Symbols(t *testing.T) {
	input := `=+(){},;`
	tests := []LexerTestCase{
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

	l := New(input)

	for i, tt := range tests {
		tt.Check(l, i, t)
	}
}

func TestLexer_Assignment(t *testing.T) {
	input := `let five = 5; let ten = 10;`

	tests := []LexerTestCase{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tt.Check(l, i, t)
	}
}

func TestLexer_WhitespaceIgnored(t *testing.T) {
	input := `

let   x    =    5    ;        
`
	tests := []LexerTestCase{
		{token.LET, "let"},
		{token.IDENT, "x"},
		{token.EQ, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tt.Check(l, i, t)
	}
}

func TestLexer_Functions(t *testing.T) {
	input := `let add = fn(x,y) { x + y; }; let result = add(five, ten);`

	tests := []LexerTestCase{
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tt.Check(l, i, t)
	}
}

func TestLexer_IfStatement(t *testing.T) {
	input := `if (5 < 10) { return true; } else { return false; }`

	tests := []LexerTestCase{
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tt.Check(l, i, t)
	}
}
func TestLexer_Operators(t *testing.T) {
	input := `!-/*5; 5 < 10 > 5; 10 == 10; 10 != 9;`

	tests := []LexerTestCase{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tt.Check(l, i, t)
	}
}
