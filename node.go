package main

type Node interface {
	Type() NodeType
}

// NodeType identifies the type of a parse tree node.
type NodeType int

// Pos represents a byte position in the original input text
type Pos int

func (p Pos) Position() Pos {
	return p
}

// Type returns itself and provides an easy default implementation
// for embedding in a Node. Embedded in all non-trivial Nodes.
func (t NodeType) Type() NodeType {
	return t
}

func (t NodeType) String() string {
	switch t {
	case NodeExpression:
		return "Expression"
	case NodeList:
		return "List"
	case NodePath:
		return "Path"
	case NodeIdentifier:
		return "Identifier"
	case NodeVariable:
		return "Variable"
	case NodeCall:
		return "Call"
	case NodeString:
		return "String"
	default:
		return "UNKNOWN NODE TYPE"
	}
}

const (
	NodeExpression NodeType = iota
	NodeList
	NodePath
	NodeIdentifier
	NodeVariable
	NodeCall
	NodeString
)

// ListNode holds a sequence of nodes.
type ListNode struct {
	NodeType
	Pos
	Nodes []Node // The element nodes in lexical order.
}

func (t *Tree) newList(pos Pos) *ListNode {
	return &ListNode{NodeType: NodeList, Pos: pos}
}

func (l *ListNode) append(n Node) {
	l.Nodes = append(l.Nodes, n)
}

type ExpressionNode struct {
	NodeType
	Pos
	Child Node
}

func (t *Tree) newExpression(pos Pos, child Node) *ExpressionNode {
	return &ExpressionNode{Pos: pos, NodeType: NodeExpression, Child: child}
}

type VariableNode struct {
	NodeType
	Pos
	Name string
}

func (t *Tree) newVariable(pos Pos, name string) *VariableNode {
	return &VariableNode{Pos: pos, Name: name, NodeType: NodeVariable}
}

type CallNode struct {
	NodeType
	Pos
	Identifier      *IdentifierNode
	Arguments       []Node
	AlternativeExpr *ExpressionNode
}

func (t *Tree) newCall(pos Pos, ident *IdentifierNode) *CallNode {
	return &CallNode{Pos: pos, Identifier: ident, NodeType: NodeCall}
}

func (n *CallNode) appendArgument(arg Node) {
	n.Arguments = append(n.Arguments, arg)
}

type PathNode struct {
	NodeType
	Pos
	Segments []Node // path segments from left to right, without separators
}

func (t *Tree) newPath(pos Pos) *PathNode {
	return &PathNode{Pos: pos, NodeType: NodePath}
}

func (n *PathNode) appendSegment(node Node) {
	n.Segments = append(n.Segments, node)
}

type IdentifierNode struct {
	NodeType
	Pos
	Value string
}

func (t *Tree) newIdentifier(pos Pos, value string) *IdentifierNode {
	return &IdentifierNode{Pos: pos, NodeType: NodeIdentifier, Value: value}
}

type StringNode struct {
	NodeType
	Pos
	Value string
}

func (t *Tree) newString(pos Pos, value string) *StringNode {
	return &StringNode{Pos: pos, NodeType: NodeString, Value: value}
}
