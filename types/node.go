package types

import (
	"fmt"
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

func (n *Node) GetValue() string {
	return n.value
}

func (n *Node) GetChildValues() MultilineValue {
	walker := &foldChildrenWalker{}
	WalkVLR(n, walker)

	return walker.Value
}

func (n *Node) GetChildValuesWithLevel(initLevel int) []LevelValuePair {
	walker := &foldLevelWalker{}
	WalkVLR(n, walker)

	return walker.Values
}
