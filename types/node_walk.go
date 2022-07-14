package types

type Walker interface {
	onValue(value string)
	beforeReturn()
}

type BaseWalker struct{}

func (w *BaseWalker) onValue(value string) {}
func (w *BaseWalker) beforeReturn()        {}

var _ Walker = &BaseWalker{}

// Base walk in order: Value - Left to Right
func WalkVLR(n *Node, walker Walker) {
	walker.onValue(n.value)

	for i := range n.children {
		child := &n.children[i]

		WalkVLR(child, walker)
	}

	walker.beforeReturn()
}
