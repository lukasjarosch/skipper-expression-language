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

const (
	NodeExpression NodeType = iota
	NodeValueExpression
	NodePathExpression
	NodeList
)

// ListNode holds a sequence of nodes.
type ListNode struct {
	NodeType
	Pos
	tr    *Tree
	Nodes []Node // The element nodes in lexical order.
}

func (t *Tree) newList(pos Pos) *ListNode {
	return &ListNode{tr: t, NodeType: NodeList, Pos: pos}
}

func (l *ListNode) append(n Node) {
	l.Nodes = append(l.Nodes, n)
}

func (l *ListNode) tree() *Tree {
	return l.tr
}

type ExpressionNode struct {
	NodeType
	Pos
}

func (t *Tree) newExpression(pos Pos) *ExpressionNode {
	return &ExpressionNode{Pos: pos, NodeType: NodeExpression}
}
