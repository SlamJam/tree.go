package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleDatanode(t *testing.T) {
	_, err := Parse(`\asdada`)
	assert.NoError(t, err)
}

func TestSingleStructnode(t *testing.T) {
	_, err := Parse(`asdada`)
	assert.NoError(t, err)
}

func TestStructnodeChain(t *testing.T) {
	_, err := Parse(`asdada foo`)
	assert.NoError(t, err)
}

func TestStructnodeChainAndDatanode(t *testing.T) {
	_, err := Parse(`asdada foo \asdada`)
	assert.NoError(t, err)
}

func TestSingleComment(t *testing.T) {
	_, err := Parse(`#asdada`)
	assert.NoError(t, err)
}

func TestEmptyLines(t *testing.T) {
	_, err := Parse(`#asdada` + "\n\n\n\n")
	assert.NoError(t, err)

	_, err = Parse("\n#     \n")
	assert.NoError(t, err)

}

func TestS(t *testing.T) {
	expression := `
bar baz
asdada foo \asdadada
	\ba		z    zz
		bazzz1 foo bar baz
			zzzz
		# nested comment
		bazzz2
		bazzz3
	level2s


multiline
	line1
	line2
`

	tree, err := Parse(expression)
	if assert.NoError(t, err) {
		_, ok := tree.Find("asdada", "foo", "asdadada", "ba		z    zz")
		assert.True(t, ok)

		node, ok := tree.Find("bar", "baz")
		if assert.True(t, ok) {
			assert.Equal(t, "baz", node.GetSelfValue())
			assert.Equal(t, "", node.GetChildFullValue())
		}

		node, ok = tree.Find("multiline")
		if assert.True(t, ok) {
			assert.Equal(t, "line1\nline2", node.GetChildFullValue())
		}

		node, ok = tree.Find("multiline2", "")
		assert.False(t, ok)
		assert.Equal(t, "", node.GetFullValue())
	}
}
