package types

import "errors"

type Line struct {
	level int
	nodes []Node
}

func NewLine(level int, nodes []Node) (Line, error) {
	if len(nodes) == 0 {
		return Line{}, errors.New("Lines with no nodes is forbidden")
	}

	return Line{level: level, nodes: nodes}, nil
}

func (l *Line) FoldToNode() Node {
	current := l.nodes[len(l.nodes)-1]
	tail := l.nodes[:len(l.nodes)-1]

	for len(tail) != 0 {
		n := tail[len(tail)-1]
		n.addChild(current)
		current = n
		tail = tail[:len(tail)-1]
	}

	return current
}

func (l *Line) FoldIntoLastNode(other Line) {
	lastNode := &l.nodes[len(l.nodes)-1]
	lastNode.addChild(other.FoldToNode())
}
