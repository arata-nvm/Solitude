package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	INT               = "INT"
	PLUS              = "PLUS"
	MINUS             = "MINUS"
	EOF               = "EOF"
)

type Token struct {
	Type    TokenType
	Literal string
}

func New(tokenType TokenType, literal string) Token {
	tok := Token{
		Type:    tokenType,
		Literal: literal,
	}

	return tok
}