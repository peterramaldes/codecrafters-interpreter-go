package main

import (
	"fmt"
	"os"
	"strings"
)

type TokenType int

const (
	EOF TokenType = iota
	LEFT_PAREN
	RIGHT_PAREN
)

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "EOF"
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	default:
		return "UNKNOWN"
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

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	var tokens []Tokenize
	if len(content) > 0 {

		// Meaning we have data on the file to read into it
		for i := 0; i < len(content); i++ {
			c := content[i]

			if len(string(c)) == 0 {
				// Skip empty line
				break
			}

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

		for _, v := range tokens {
			fmt.Printf("%v\n", v)
		}
	} else {
		//fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
