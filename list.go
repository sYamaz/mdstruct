package mdstruct

import "strings"

type (
	MDList struct {
		listItems []MDListItem
	}

	MDListItem struct {
		Ordered     bool
		IndentLevel int
		MDInlineArray
	}
)

func NewMDList(items []MDListItem) *MDList {
	i := items
	if i == nil {
		i = []MDListItem{}
	}

	return &MDList{
		listItems: i,
	}
}

func (md *MDList) ToMDBlockString() string {
	strs := []string{}
	for _, item := range md.listItems {
		strs = append(strs, item.ToString())
	}
	return strings.Join(strs, "\n")
}

func (md *MDList) ToMDString() string {
	return md.ToMDBlockString() + "\n"
}

func (md *MDList) AddListItem(ordered bool, indentLevel int, inlines MDInlineArray) *MDList {
	md.listItems = append(md.listItems, MDListItem{
		Ordered:       ordered,
		IndentLevel:   indentLevel,
		MDInlineArray: inlines,
	})

	return md
}

func (md *MDListItem) ToString() string {
	str := joinMDInlines(md.MDInlineArray)
	mark := "- "
	if md.Ordered {
		mark = "1. "
	}

	indent := strings.Repeat("  ", md.IndentLevel)

	return indent + mark + str
}
