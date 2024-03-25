package main

import (
	"fmt"
)

type Tree struct {
	root      *ListNode
	lex       *lexer
	input     string
	token     [3]Token // lookahead buffer
	peekCount int
}

func Parse(text string) (*Tree, error) {
	t := &Tree{}
	return t.Parse(text)
}

func (t *Tree) Parse(input string) (*Tree, error) {
	t.input = input
	t.lex = lex(input)
	t.parse()

	return nil, nil
}

// expect consumes the next token and guarantees it has the required type.
func (t *Tree) expect(expected TokenType, context string) Token {
	token := t.next()
	if token.Type != expected {
		t.unexpected(token, context)
	}
	return token
}

// unexpected complains about the token and terminates processing.
func (t *Tree) unexpected(token Token, context string) {
	if token.Type == tError {
		t.errorf("%s in %s: %s", TokenString(token.Type), context, token.Value)
	}
	t.errorf("unexpected %s in %s", TokenString(token.Type), context)
}

// errorf formats the error and terminates processing.
func (t *Tree) errorf(format string, args ...any) {
	t.root = nil
	panic(fmt.Errorf(format, args...))
}

func (t *Tree) peek() Token {
	if t.peekCount == 2 {
		return t.token[t.peekCount-2]
	}
	if t.peekCount == 1 {
		return t.token[t.peekCount-1]
	}
	t.peekCount = 1
	t.token[0] = t.lex.nextToken()
	return t.token[0]
}

// peek2 returns but does not consume the next two tokens.
func (t *Tree) peek2() (Token, Token) {
	if t.peekCount == 0 {
		t.peekCount = 1
		t.token[0] = t.lex.nextToken()
	}
	if t.peekCount == 1 {
		t.peekCount = 2
		t.token[1] = t.lex.nextToken()
	}
	if t.peekCount == 2 {
		return t.token[0], t.token[1]
	}
	return Token{}, Token{}
}

// backup backs the input stream up one token.
func (t *Tree) backup() {
	t.peekCount++
}

// // backup2 backs the input stream up by two tokens.
// func (t *Tree) backup2(token1, token2 Token) {
// 	t.token[0] = token1
// 	t.token[1] = token2
// 	t.peekCount = 2
// }

// next returns the next token.
func (t *Tree) next() Token {
	// if t.peekCount == 2 {
	// 	t.peekCount -= 2
	// 	return t.token[t.peekCount]
	// }
	// if t.peekCount == 1 {
	// 	t.peekCount--
	// 	return t.token[t.peekCount]
	// }
	// t.peekCount = 1
	// t.token[0] = t.lex.nextToken()
	// return t.token[0]
	fmt.Println(t.peekCount, t.token)
	if t.peekCount == 2 {
		t.peekCount--
		t.peekCount--
		return t.token[t.peekCount]
	} else if t.peekCount == 1 {
		t.peekCount--
		return t.token[t.peekCount]
	} else {
		t.token[0] = t.lex.nextToken()
	}
	return t.token[t.peekCount]
}

func (t *Tree) parse() {
	// t.root = t.newList(Pos(t.peek().Pos))
	for t.peek().Type != tEOF {
		switch tok := t.next(); tok.Type {
		case tLeftDelim:
			n := t.parseExpression()
			t.root.append(n)
		default:
			t.unexpected(tok, "parse")
		}

		// n := t.parseExpression()
		// t.root.append(n)

		// // the first token must be a left delimiter to start an expression
		// // all other tokens are dropped
		// if t.peek().Type == tLeftDelim {
		// 	t.next() // discard delimiter
		//
		// 	tok := t.peek().Type
		// 	switch tok {
		// 	case tIdent:
		// 		ident := t.next()
		// 		fmt.Println("IDENT", ident)
		// 	case tLeftDelim:
		// 		t.next()
		// 		fmt.Println("LEFT DELIM")
		// 	default:
		// 		t.errorf("syntax error: unexpected token within expression")
		// 		return
		// 	}
		// }
		//
		// // TODO: what about errors? Can they occur here?
		// t.next() // discard token, we're not within an expression
	}
	t.errorf("unexpected EOF")
}

// parseExpression:
//
//	expression = (value_expression | path_expression) ;
//
// The left delimiter is already consumed at this point.
func (t *Tree) parseExpression() Node {
	// To distinguish between path_expression and value_expression
	// we need to perform a lookahead of 2.
	// Each expression starts with an identifier (t1),
	// and the path_expression must have a tPathSep after that (t2)
	t1, t2 := t.peek2()
	if t1.Type != tIdent {
		t.unexpected(t1, "parseExpression")
	}

	// the second token is tPathSep -> parse path_expression
	if t2.Type == tPathSep {
		return t.parsePathExpression()
	}

	// otherwise parse value_expression
	return t.parseValueExpression()
}

// parsePathExpression
func (t *Tree) parsePathExpression() Node {
	context := "parsePathExpression"

	fmt.Println(t.next())
	fmt.Println(t.next())
	fmt.Println(t.next())
	fmt.Println(t.next())
	fmt.Println(t.next())
	// t.expect(tIdent, context)
	// t.expect(tPathSep, context)
	return nil

	for {
		switch tok := t.peek(); tok.Type {
		case tIdent:
			fmt.Println("IDENT", tok)
			t.next()
		case tPathSep:
			fmt.Println("SEP")
			t.next()
		case tRightDelim:
			fmt.Println("END")
			t.next()
		case tLeftDelim:
			fmt.Println("START NEW")
			t.next()
		case tEOF:
			fallthrough
		default:
			t.unexpected(tok, context)
			return nil
		}
	}
}

func (t *Tree) parseValueExpression() Node {
	t.unexpected(t.next(), "OASDF")
	return nil
}
