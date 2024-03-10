# mdstruct

mdstructパッケージはgolangのstructでmarkdownを表現し、markdownフォーマットのテキストとして出力するためのパッケージです。

## usage

[example](./example/create-document/main.go)

```go
package main

import (
	"log"

	"github.com/sYamaz/mdstruct"
)

func main() {
	document := mdstruct.NewMDDocument(nil)

	document.
		Add(mdstruct.NewMDHeadline(1, mdstruct.NewMDText("headline1"))).
		Add(mdstruct.NewMDText("text.").
            Then(mdstruct.NewMDCode(mdstruct.NewMDText("code")).
            Then(mdstruct.NewMDBold(mdstruct.NewMDText("bold"))))).
		Add(mdstruct.NewMDHeadline(2, mdstruct.NewMDText("headline2"))).
		Add(mdstruct.NewMDList(nil).
			AddListItem(false, 0, mdstruct.NewMDText("item1")).
			AddListItem(false, 0, mdstruct.NewMDText("item2")).
			AddListItem(false, 0, mdstruct.NewMDText("item3")))

	log.SetFlags(0)
	log.SetPrefix("")
	log.Println(document.ToMDString())
}

// Output
// # headline1
//
// text.`code`**bold**
//
// ## headline2
//
// - item1
// - item2
// - item3
```

## support

||support|
|---|---|
|code span|not tested|
|code block|-|
|bold(Emphasis)|not tested|
|italic|not tested|
|strikethrough|-|
|ordered list|not tested|
|unordered list|not tested|
|checkbox list|-|
|table|not tested|
|headline|not tested|
|blockquotes|-|
|horizontal rules|-|
|link|-|
|img|-|