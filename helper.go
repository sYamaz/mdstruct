package mdstruct

import "strings"

func joinMDInlines(inlines MDInlineArray) string {
	if inlines == nil {
		return ""
	}

	strs := []string{}
	for _, i := range inlines {
		strs = append(strs, i.ToInlineString())
	}
	return strings.Join(strs, "")
}
