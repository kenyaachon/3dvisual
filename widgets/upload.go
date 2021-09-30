package widgets

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
)

// UploadWidget implements widget responsible for uploading JSON file.
type UploadWidget struct {
	vecty.Core

	handler func([]byte)
}

// NewUploadWidget creates a new upload widget.
func NewUploadWidget(handler func([]byte)) *UploadWidget {
	return &UploadWidget{
		handler: handler,
	}
}

func (u *UploadWidget) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("file", "is-primary", "is-small"),
		),
		elem.Label(
			vecty.Markup(
				vecty.Class("file-label"),
			),
			elem.Input(
				vecty.Markup(
					vecty.Class("file-input"),
					prop.ID("file"),
					prop.Type("file"),
					vecty.Property("accept", "application/json"), // TODO(divan): add prop.Accept PR
					event.Input(u.onUploadClick),
				),
			),
			elem.Span(
				vecty.Markup(
					vecty.Class("file-cta"),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("file-icon"),
					),
					elem.Italic(
						vecty.Markup(
							vecty.Class("fas", "fa-upload"),
						),
					),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("file-label"),
					),
					vecty.Text("Upload..."),
				),
			),
		),
	)
}

// onUploadClick implements callback for "Upload" button clicked event.
func (u *UploadWidget) onUploadClick(e *vecty.Event) {
	// FIXME: run as a gorotine because GopherJS can't block JS thread in callback
	// go func() {
	// 	file := e.Target.Get("files").Index(0)
	// 	fmt.Println("File size:", file.Get("size"))

	// 	data := jsapi.NewFileReader().ReadAll(file)
	// 	u.handler(data)
	// }()
}