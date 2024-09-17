package scan

import (
	"fmt"
	"strings"
)

// TODO: it's better to use rune here?
type TokenType int

const (
	EOF TokenType = iota
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	STAR
	DOT
	COMMA
	PLUS
	MINUS
	SEMICOLON
)

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "EOF"
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case STAR:
		return "STAR"
	case DOT:
		return "DOT"
	case COMMA:
		return "COMMA"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case SEMICOLON:
		return "SEMICOLON"
	default:
		return "UNKOWN"
	}
}

type Tokenize struct {
	tokenType TokenType

	// The actual sequence of characters that formed the token
	// For an EOF token, the lexeme is an empty string.
	lexeme string

	// The literal value of the token
	// For most tokens this is null.
	// For `STRING` and `NUMBER` tokens, it holds the value of the string/number.
	literal string
}

func (t Tokenize) String() string {
	return fmt.Sprintf("%s %s %s", t.tokenType, t.lexeme, t.literal)
}

func Scan(content []byte) []Tokenize {
	var tokens []Tokenize

	// Meaning we have data on the file to read into it
	for i := 0; i < len(content); i++ {
		c := content[i]

		if c == '\n' {
			// Skip breakline
			break
		}

		// Parse token type
		var tt TokenType
		switch string(c) {
		case "(":
			tt = LEFT_PAREN
		case ")":
			tt = RIGHT_PAREN
		case "{":
			tt = LEFT_BRACE
		case "}":
			tt = RIGHT_BRACE
		case "*":
			tt = STAR
		case ".":
			tt = DOT
		case ",":
			tt = COMMA
		case "+":
			tt = PLUS
		case "-":
			tt = MINUS
		case ";":
			tt = SEMICOLON
		default:
			panic(fmt.Sprintf("`%+v` not implemented yet", string(c)))
		}

		// Parse the lexeme
		lx := strings.TrimSpace(string(c))

		// Parse literal
		l := "null"

		val := Tokenize{
			tokenType: tt,
			lexeme:    lx,
			literal:   l,
		}

		tokens = append(tokens, val)
	}

	// Add `EOF null` as we finish the file
	eof := Tokenize{
		tokenType: EOF,
		lexeme:    "",
		literal:   "null",
	}

	tokens = append(tokens, eof)

	return tokens
}
