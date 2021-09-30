package widgets

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
)

// ForceInput represents input widget for forces.
type ForceInput struct {
	vecty.Core

	changed     bool
	title       string
	description string
	value       float64
}

// NewForceInput creates new input.
func NewForceInput(title, description string, value float64) *ForceInput {
	f := &ForceInput{
		title:       title,
		description: description,
		value:       value,
	}
	return f
}

// Render implements vecty.Component interface for ForceInput.
func (f *ForceInput) Render() vecty.ComponentOrHTML {
	value := fmt.Sprintf("%.4f", f.value)
	return InputField(f.title, f.description,
		elem.Input(
			vecty.Markup(
				vecty.Class("input", "is-small"),
				vecty.Style("text-align", "right"),
				prop.Value(value),
				event.Input(f.onEditInput),
			),
		),
	)
}

// helper for wrapping many divs
func fieldControl(element vecty.MarkupOrChild) *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("field-body"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("field"),
			),
			elem.Paragraph(
				vecty.Markup(
					vecty.Class("control"),
				),

				element,
			),
		),
	)
}

func (f *ForceInput) onEditInput(event *vecty.Event) {
	value := event.Target.Get("value").Float()

	f.changed = true
	f.value = value
}

// Value returns the current value.
func (f *ForceInput) Value() float64 {
	return f.value
}

// Changed returns if input value has been changed, and resets it's value to false.
func (f *ForceInput) Changed() bool {
	if f.changed {
		f.changed = false
		return true
	}

	return false
}
