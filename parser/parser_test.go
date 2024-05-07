package parser

import (
	"gorangutang/ast"
	"gorangutang/lexer"
	"testing"
)

func TestParser_LetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf(
			"program.Statements does not contain 3 statements. got=%d",
			len(program.Statements),
		)
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors:", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letSt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not *ast.LetStatement. got=%T", s)
		return false
	}

	if letSt.Name.Value != name {
		t.Errorf("letSt.Name.Value not %s. got=%s", name, letSt.Name.Value)
		return false
	}

	if letSt.Name.TokenLiteral() != name {
		t.Errorf(
			"letSt.Name.TokenLiteral() not %s. got=%s",
			letSt.Name.TokenLiteral(),
			name,
		)
		return false
	}

	return true
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return a;
	return add(5);
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements != 3. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		retStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Statement not *ast.ReturnStatement. got=%T", stmt)
			continue
		}
		if retStmt.TokenLiteral() != "return" {
			t.Errorf("Statement's Token Literal was not token.RETURN. got=%q", retStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := `foobar;`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements != 1. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("prgram.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("stmt.Expression not *ast.Identifier. got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not \"foobar\". got=%s", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral() not \"foobar\". got=%s", ident.TokenLiteral())
	}
}
