package mdstruct

type (
	MDInlineArray []MDInline
)

func (md MDInlineArray) Then(inline MDInlineArray) MDInlineArray {
	ret := []MDInline{}

	ret = append(ret, md...)
	ret = append(ret, inline...)

	return ret
}

func (md MDInlineArray) ToMDString() string {
	return joinMDInlines(md)
}
