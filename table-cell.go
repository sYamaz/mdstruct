// table-cell is not markdown element
// it is only table content
package mdstruct

type (
	MDTableCell struct {
		contents MDInlineArray
	}
)

func NewMDTableCell(contents MDInlineArray) *MDTableCell {
	c := contents
	if c == nil {
		c = []MDInline{}
	}

	return &MDTableCell{
		contents: c,
	}
}

func (md *MDTableCell) toInlineString() string {
	return joinMDInlines(md.contents)
}

func (md *MDTableCell) toString() string {
	return md.toInlineString()
}
