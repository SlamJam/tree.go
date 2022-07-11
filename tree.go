package tree

import (
	tree "github.com/SlamJam/tree.go/internal"
	"github.com/SlamJam/tree.go/types"
)

func Parse(input string) (types.Tree, error) {
	res, err := tree.Parse("input", []byte(input))
	if err != nil {
		return types.Tree{}, err
	}

	return res.(types.Tree), nil
}
