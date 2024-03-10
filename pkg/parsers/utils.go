package estropadakParser

import (
	"strings"

	"golang.org/x/net/html"
)

// ByPosition implements sort.Interface for []Result based on
// the Position field.
type ByPosition []Result

func (a ByPosition) Len() int           { return len(a) }
func (a ByPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPosition) Less(i, j int) bool { return a[i].Position < a[j].Position }

// ByName implements sort.Interface for []Result based on
// the Name field.
type ByName []Result

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].TeamName < a[j].TeamName }

// ByTime implements sort.Interface for []Result based on
// the Time field.
type ByTime []Result

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return a[i].Time < a[j].Time }

func attr_has_value(t html.Tokenizer, attrName string, value string) bool {
	result := false
	for {
		attr, val, more_attrs := t.TagAttr()
		if string(attr) == attrName && strings.Contains(string(val), value) {
			result = true
			break
		}
		if !more_attrs {
			break
		}
	}
	return result
}
