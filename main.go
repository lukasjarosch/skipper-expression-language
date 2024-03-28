package main

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	input := `
	thjere can be very much random noisy text, we don't care
	foo: bar
	| asdf
	${FOO}
	ohai ${$name}
	Hello, ${foo:bar:baz} more
	${call()}
	${foo:$target:baz}
	text ${VARIABLE_NAME}
	${reveal(foo:bar:baz, another_call())}
	${call($target_name) || env("FOO")}
	${call('param ', "another very long string   ")}
	${reveal(secrets:$target_name:foobar) || env(config:secrets:default) }
	${reveal($secrets,$target_name,$password) || env(config:secrets:default)}
	${foo}
`

	input = `${foo("foo", $bar, call('foo') || foo())}`
	// input = `${call($a)}`
	// l := lex(input)
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

	// scanner := sel.NewScanner(input, os.Stderr)
	// tokens := scanner.ScanTokens()
	//
	// for _, tok := range tokens {
	// 	fmt.Println(tok)
	// }

	ast, _ := Parse(input)
	// spew.Dump(err)

	for _, expr := range ast.Nodes {
		spew.Println(expr.Type())
		child := expr.(*ExpressionNode).Child

		switch child.Type() {
		case NodeVariable:
			varNode := child.(*VariableNode)
			spew.Printf(" └ Variable: %s\n", varNode.Name)

		case NodePath:
			pathNode := child.(*PathNode)
			spew.Printf(" └ Path: %s\n", pathNode.Segments)

		case NodeCall:
			callNode := child.(*CallNode)
			args := []string{}

			for _, arg := range callNode.Arguments {
				args = append(args, arg.Type().String())
			}

			spew.Printf(" └ Call: %s(%s)\n", callNode.Identifier.Value, strings.Join(args, ","))
		}

	}
}
