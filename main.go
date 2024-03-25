package main

import (
	"fmt"
)

func main() {
	input := `
	foo: bar
	Hello, ${foo:bar:baz} more
	text ${VARIABLE_NAME}
	${foo:${bar:baz}}
	${call(${variable}) || env(FOO)}
	${call('param ', "another very long string   ")}
	${reveal(secrets:${target_name}:password) || env(config:secrets:default) }
	${reveal(${secrets:${target_name}:password}) || env(${config:secrets:default}) }
`

	// l := lex(input)
	// for t := range l.tokens {
	// 	if t.Type == tString {
	// 		fmt.Printf("|%3d| %-20s | \"%s\"\n", t.Pos, TokenString(t.Type), t.Value)
	// 		// continue
	// 		break
	// 	}
	// 	fmt.Printf("|%3d| %-20s | %s\n", t.Pos, TokenString(t.Type), t.Value)
	// }

	// for {
	// 	t := l.nextToken()
	// 	if t.Type == tString {
	// 		fmt.Printf("|%3d| %-20s | \"%s\"\n", t.Pos, TokenString(t.Type), t.Value)
	// 		continue
	// 	}
	// 	fmt.Printf("|%3d| %-20s | %s\n", t.Pos, TokenString(t.Type), t.Value)
	//
	// 	if t.Type == tEOF || t.Type == tError {
	// 		break
	// 	}
	// }

	ast, err := Parse(input)
	fmt.Println(err)
	fmt.Println(ast)
}
