package types

import (
	"errors"
	"fmt"
)

type Tree struct {
	nodes []Node
}

func (t *Tree) Find(path ...string) (Node, bool) {
	nodes := t.nodes
	var result *Node

loop:
	for i := 0; i < len(path); i++ {
		name := path[i]

		for j := 0; j < len(nodes); j++ {
			node := &nodes[j]
			if node.value == name {
				result = node
				nodes = node.children
				continue loop
			}
		}

		result = nil
		break
	}

	if result == nil {
		return Node{}, false
	}

	return *result, true
}

func NewTree(lines []Line) (Tree, error) {
	var prevLevel int
	for i := range lines {
		line := &lines[i]

		if i == 0 {
			if line.level != 0 {
				return Tree{}, fmt.Errorf("first line always must be at level 0")
			}
		} else {
			if line.level-prevLevel > 1 {
				return Tree{}, fmt.Errorf("line indent may be %d or less, not %d", prevLevel+1, line.level)
			}
		}

		prevLevel = line.level
	}

	revLines := make([]Line, len(lines))
	copy(revLines, lines)

	for i, j := 0, len(revLines)-1; i < j; i, j = i+1, j-1 {
		revLines[i], revLines[j] = revLines[j], revLines[i]
	}

	var roots []Node

	for len(revLines) != 0 {
		prev := revLines[0]
		tail := revLines[1:]

		if prev.level == 0 {
			roots = append(roots, prev.FoldToNode())
			revLines = tail
			continue
		}

		var deffered []Line
		for i := range tail {
			current := &tail[i]

			if current.level >= prev.level {
				deffered = append(deffered, prev)
			} else {
				current.FoldIntoLastNode(prev)

				for len(deffered) != 0 {
					lastDeffered := deffered[len(deffered)-1]
					if lastDeffered.level <= current.level {
						break
					}

					current.FoldIntoLastNode(lastDeffered)
					deffered = deffered[:len(deffered)-1]
				}

			}

			if current.level == 0 {
				if len(deffered) != 0 {
					return Tree{}, errors.New("internal error: deffered is not empty")
				}

				roots = append(roots, current.FoldToNode())
				start := i + 1
				if start > len(tail) {
					start = len(tail)
				}
				revLines = tail[start:]
				break
			}

			prev = *current
		}
	}

	return Tree{nodes: roots}, nil
}
