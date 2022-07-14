package types

import "strings"

/* */

type MultilineValue []string

func (v MultilineValue) String() string {
	return strings.Join(v, "\n")
}

type foldChildrenWalker struct {
	BaseWalker
	level int

	Value MultilineValue
}

func (w *foldChildrenWalker) onValue(value string) {

	if w.level != 0 {
		w.Value = append(w.Value, value)
	}

	w.level++
}
func (w *foldChildrenWalker) beforeReturn() { w.level-- }

/* */

type LevelValuePair struct {
	Level int
	Value string
}
type foldLevelWalker struct {
	BaseWalker
	level int

	Values []LevelValuePair
}

func (w *foldLevelWalker) onValue(value string) {
	w.Values = append(w.Values, LevelValuePair{Level: w.level, Value: value})
	w.level++
}
func (w *foldLevelWalker) beforeReturn() { w.level-- }
