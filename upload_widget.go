package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
)

// UploadWidget implements widget responsible for uploading JSON file.
// TODO(divan): move to widgets package
type UploadWidget struct {
	vecty.Core

	handler func(e *vecty.Event)
}

// NewUploadWidget creates a new upload widget.
func NewUploadWidget(handler func(*vecty.Event)) *UploadWidget {
	return &UploadWidget{
		handler: handler,
	}
}

func (u *UploadWidget) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Input(
			vecty.Markup(
				prop.ID("file"),
				prop.Type("file"),
				vecty.Property("accept", "application/json"), // TODO(divan): add prop.Accept PR
				event.Input(u.handler),
			),
			vecty.Text("Upload network.json"),
		),
		elem.Span(
			vecty.Text("Using graph with"),
		),
	)
}
