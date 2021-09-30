package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

func Tile(content vecty.MarkupOrChild) *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("tile"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("tile", "is-child", "box"),
			),
			content,
		),
	)
}

func TileParent(content vecty.MarkupOrChild) *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("tile", "is-parent"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("tile", "is-child", "box"),
			),
			content,
		),
	)
}
