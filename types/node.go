package types

import (
	"fmt"
	"strings"
)

type Node struct {
	is_data bool
	value   string

	children []Node
}

func NewStructNode(value string) Node {
	return Node{
		value: value,
	}
}

func NewDataNode(value string) Node {
	return Node{
		is_data: true,
		value:   value,
	}
}

func (n Node) String() string {
	if n.is_data {
		return fmt.Sprintf("DataNode(%s)", n.value)
	}

	return fmt.Sprintf("StructNode(%s)", n.value)
}

func (n *Node) addChild(ch Node) {
	n.children = append(n.children, ch)
}

func (n *Node) GetSelfValue() string {
	return n.value
}

func (n *Node) GetFullValue() string {
	return strings.Join(n.GetFullValueMultiline(), "\n")
}

func (n *Node) GetFullValueMultiline() []string {
	values := make([]string, 0, len(n.children))
	values = append(values, n.value)

	for i := range n.children {
		node := &n.children[i]

		values = append(values, node.GetFullValueMultiline()...)
	}

	return values
}

func (n *Node) GetChildFullValue() string {
	return strings.Join(n.GetChildFullValueMultiline(), "\n")
}

func (n *Node) GetChildFullValueMultiline() []string {
	values := make([]string, 0, len(n.children))

	for i := range n.children {
		node := &n.children[i]

		values = append(values, node.GetFullValueMultiline()...)
	}

	return values
}

type LeveledValue struct {
	Level int
	Value string
}

func (n *Node) GetChildValueWithLevels(initLevel int) []LeveledValue {
	values := make([]LeveledValue, 0, len(n.children))

	for i := range n.children {
		node := &n.children[i]
		values = append(values, LeveledValue{Level: initLevel, Value: node.value})
		values = append(values, node.GetChildValueWithLevels(initLevel+1)...)
	}

	return values
}
