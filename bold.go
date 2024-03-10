package mdstruct

type (
	MDBold struct {
		contents []MDInline
	}
)

func NewMDBold(contents []MDInline) MDInlineArray {
	c := contents
	if c == nil {
		c = []MDInline{}
	}
	return []MDInline{&MDBold{
		contents: c,
	}}
}

func (md *MDBold) ToInlineString() string {
	return "**" + joinMDInlines(md.contents) + "**"
}

// implement MDElement
func (md *MDBold) ToMDString() string {
	return md.ToInlineString()
}

// Then implements MDInline.
func (md *MDBold) Then(inline MDInline) MDInlineArray {
	ret := MDInlineArray{}
	ret = append(ret, md, inline)
	return ret
}

func (md *MDBold) ToInlines() MDInlineArray {
	return []MDInline{md}
}
