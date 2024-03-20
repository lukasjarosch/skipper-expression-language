package main

import (
	"fmt"
)

func main() {
	input := `
	Hello, ${foo:bar:baz} more
	text ${VARIABLE_NAME}
	${foo:${bar:baz}}
	${call(${variable}) || env(FOO)}
	${call('param ', "another very long string   ")}
	${reveal(secrets:${target_name}:password) || env(config:secrets:default) }
`

	l := lex(input)
	for t := range l.tokens {
		if t.Type == tString {
			fmt.Printf("|%3d| %-20s | \"%s\"\n", t.Pos, TokenString(t.Type), t.Value)
			continue
		}
		fmt.Printf("|%3d| %-20s | %s\n", t.Pos, TokenString(t.Type), t.Value)
	}
}
