//go:generate go run github.com/mna/pigeon@latest -optimize-parser -o tree.peg.go tree.peg

package tree

import (
	"errors"
	"fmt"

	"github.com/SlamJam/tree.go/types"
)

func parseLine(tabs []interface{}, items []interface{}, prev interface{}, on_level func(level int)) (types.Line, error) {
	if len(items) != 2 {
		return types.Line{}, errors.New("items must be length 2")
	}

	pairs := items[0].([]interface{})
	nodes := make([]types.Node, 0, len(pairs)+1)

	for i := range pairs {
		pair := pairs[i].([]interface{})

		if len(items) != 2 {
			return types.Line{}, fmt.Errorf("(StructNode, space) pair must be length 2: %s", pair)
		}

		struct_node := pair[0].(types.Node)
		nodes = append(nodes, struct_node)
	}

	if data_node, ok := items[1].(types.Node); ok {
		nodes = append(nodes, data_node)
	}

	indent := len(tabs)

	// check line level for pretty parsing errors (on right line and pos)
	if prev == nil {
		if indent != 0 {
			return types.Line{}, fmt.Errorf("root line must be at level 0")
		}
	} else {
		prev_int := prev.(int)
		if indent-prev_int > 1 {
			return types.Line{}, fmt.Errorf("line indent may be %d or less, not %d", prev_int+1, indent)
		}
	}
	on_level(indent)

	return types.NewLine(indent, nodes)
}

func parseTree(items []interface{}) (types.Tree, error) {
	lines := make([]types.Line, 0, len(items))

	for i := range items {
		if line, ok := items[i].(types.Line); ok {
			lines = append(lines, line)
		}
	}

	return types.NewTree(lines)
}

func parseStructnode(value string) (types.Node, error) {
	return types.NewStructNode(value), nil
}

func parseDatanode(value string) (types.Node, error) {
	return types.NewDataNode(value), nil
}
