package main

import (
	"fmt"
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

	spew.Println(input)
	for _, expr := range ast.Nodes {
		spew.Println(expr.Type())
		child := expr.(*ExpressionNode).Child

		printTree(child, 0)
	}
}

// func printNode(node Node) {
// 	switch node.Type() {
// 	case NodeVariable:
// 		varNode := node.(*VariableNode)
// 		spew.Printf(" └ Variable: %s\n", varNode.Name)
//
// 	case NodePath:
// 		pathNode := node.(*PathNode)
// 		spew.Printf(" └ Path: %s\n", pathNode.Segments)
//
// 	case NodeCall:
// 		callNode := node.(*CallNode)
// 		args := []string{}
//
// 		for _, arg := range callNode.Arguments {
// 			args = append(args, arg.Type().String())
// 		}
//
// 		spew.Printf(" └ Call: %s(%s)\n", callNode.Identifier.Value, strings.Join(args, ","))
// 	}
// }

func printTree(node Node, indent int) {
	// Print the current node
	prefix := strings.Repeat("  ", indent)
	switch n := node.(type) {
	case *VariableNode:
		fmt.Printf("%s%s: %s\n", prefix, n, n.Name)
	case *IdentifierNode:
		fmt.Printf("%s%s: %s\n", prefix, n, n.Value)
	case *StringNode:
		fmt.Printf("%s%s: %q\n", prefix, n, n.Value)
	default:
		fmt.Printf("%s%s\n", prefix, node)
	}

	// Determine the type of the node and handle accordingly
	switch n := node.(type) {
	case *ListNode:
		for _, child := range n.Nodes {
			printTree(child, indent+1)
		}
	case *ExpressionNode:
		printTree(n.Child, indent+1)
	case *CallNode:
		printTree(n.Identifier, indent+1)
		for _, arg := range n.Arguments {
			printTree(arg, indent+1)
		}
		if n.AlternativeExpr != nil {
			printTree(n.AlternativeExpr, indent+1)
		}
	case *PathNode:
		for _, seg := range n.Segments {
			printTree(seg, indent+1)
		}
		// For other nodes like IdentifierNode, StringNode, VariableNode, etc.,
		// there are no children to print, so we do nothing more.
	}
}
