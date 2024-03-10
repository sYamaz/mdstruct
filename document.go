package mdstruct

import "strings"

type (
	MDDocument struct {
		contents []MDElement
	}

	MDInline interface {
		ToInlineString() string
		// implement MDElement
		ToMDString() string
		ToInlines() MDInlineArray
		Then(inline MDInline) MDInlineArray
	}

	MDBlock interface {
		ToBlockString() string
		// implement MDElement
		ToMDString() string
	}

	MDElement interface {
		ToMDString() string
	}
)

func NewMDDocument(content []MDElement) *MDDocument {
	c := content
	if c == nil {
		c = []MDElement{}
	}

	return &MDDocument{
		contents: c,
	}
}

func (md *MDDocument) Add(ele MDElement) *MDDocument {
	md.contents = append(md.contents, ele)
	return md
}

func (md *MDDocument) ToMDString() string {
	strs := []string{}

	for _, c := range md.contents {
		strs = append(strs, c.ToMDString())
	}

	return strings.Join(strs, "\n\n")
}
