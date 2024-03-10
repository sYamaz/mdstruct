package main

import (
	"log"

	"github.com/sYamaz/mdstruct"
)

func main() {
	document := mdstruct.NewMDDocument(nil)

	document.
		Add(mdstruct.NewMDHeadline(1, mdstruct.NewMDText("headline1"))).                                                                            // # headline1
		Add(mdstruct.NewMDText("text.").Then(mdstruct.NewMDCode(mdstruct.NewMDText("code")).Then(mdstruct.NewMDBold(mdstruct.NewMDText("bold"))))). // text.`code`**bold**
		Add(mdstruct.NewMDHeadline(2, mdstruct.NewMDText("headline2"))).                                                                            // ## headline2
		Add(mdstruct.NewMDList(nil).
			AddListItem(false, 0, mdstruct.NewMDText("item1")).
			AddListItem(false, 0, mdstruct.NewMDText("item2")).
			AddListItem(false, 0, mdstruct.NewMDText("item3")))

	log.SetFlags(0)
	log.SetPrefix("")
	log.Println(document.ToMDString())
}
