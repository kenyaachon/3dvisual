package widgets

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
)

// Switch is a wrapper around checkbox-like switch.
type Switch struct {
	vecty.Core

	title     string
	isChecked bool
	handler   func()

	domID string
}

// NewSwitch creates and inits a new switch.
func NewSwitch(title string, checked bool, handler func()) *Switch {
	rnd := rand.Int63n(math.MaxInt64)
	domID := fmt.Sprintf("idSwitch%d", rnd)
	return &Switch{
		title:     title,
		isChecked: checked,
		handler:   handler,
		domID:     domID,
	}
}

// Render implements vecty's Component interface for Switch.
func (s *Switch) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Input(
			vecty.Markup(
				vecty.Class("switch", "is-rounded", "is-small"),
				prop.ID(s.domID),
				prop.Type(prop.TypeCheckbox),
				event.Change(s.onToggle),
				vecty.MarkupIf(s.isChecked,
					vecty.Attribute("checked", "checked"),
				),
			),
		),
		elem.Label(
			vecty.Markup(
				// this is needed to properly handle click
				// see https://wikiki.github.io/form/switch/
				vecty.Attribute("for", s.domID),
			),
			vecty.Text(s.title),
		),
	)
}

// Checked returns the checked state of the switch.
func (s *Switch) Checked() bool {
	return s.isChecked
}

// onToggle changes the checked state of the switch.
func (s *Switch) onToggle(*vecty.Event) {
	s.isChecked = !s.isChecked
	vecty.Rerender(s)
	if s.handler != nil {
		go s.handler()
	}
}
