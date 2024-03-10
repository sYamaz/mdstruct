package mdstruct

type (
	MDCode struct {
		contents []MDInline
	}
)

func NewMDCode(contents []MDInline) MDInlineArray {
	c := contents
	if c == nil {
		c = []MDInline{}
	}
	return []MDInline{&MDCode{
		contents: c,
	}}
}

func (md *MDCode) ToInlineString() string {
	return "`" + joinMDInlines(md.contents) + "`"
}

// implement MDElement
func (md *MDCode) ToMDString() string {
	return md.ToInlineString()
}

// Then implements MDInline.
func (md *MDCode) Then(inline MDInline) MDInlineArray {
	ret := MDInlineArray{}
	ret = append(ret, md, inline)
	return ret
}

func (md *MDCode) ToInlines() MDInlineArray {
	return []MDInline{md}
}
