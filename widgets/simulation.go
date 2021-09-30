package widgets

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
)

const (
	DefaultTTL = 10 // seconds
)

// Simulation represents configuration panel for propagation simulation.
type Simulation struct {
	vecty.Core
	startSimulation func() (int, error)
	replay          func()
	step            func(int)

	address  string // backend host address
	timeline *Range

	ttl int

	errMsg     string
	hasResults bool
	inProgress bool
}

// NewSimulation creates new simulation configuration panel. If simulation
// backend host address is not specified, it'll use 'localhost:8084' as a default.
func NewSimulation(address string, startSimulation func() (int, error), replay func(), step func(int)) *Simulation {
	if address == "" {
		address = "http://localhost:8084"
	}
	return &Simulation{
		ttl:             DefaultTTL,
		address:         address,
		startSimulation: startSimulation,
		replay:          replay,
		step:            step,
	}
}

// Render implements vecty.Component interface for Simulation.
func (s *Simulation) Render() vecty.ComponentOrHTML {
	return Widget(
		elem.Div(
			Header("Simulation backend:"),
			InputField("Host:", "Simulation backend host address",
				elem.Input(
					vecty.Markup(
						vecty.MarkupIf(s.inProgress,
							vecty.Attribute("disabled", "true"),
						),
						vecty.Class("input", "is-small"),
						prop.Value(s.address),
						event.Input(s.onAddressChange),
					),
				),
			),
			InputField("TTL:", "Message time to live value (in seconds)",
				elem.Input(
					vecty.Markup(
						vecty.MarkupIf(s.inProgress,
							vecty.Attribute("disabled", "true"),
						),
						vecty.Class("input", "is-small"),
						prop.Value(fmt.Sprint(s.ttl)),
						event.Input(s.onTTLChange),
					),
				),
			),
		),
		elem.Div(
			elem.Button(
				vecty.Markup(
					vecty.MarkupIf(s.inProgress,
						vecty.Attribute("disabled", "true"),
						vecty.Class("is-loading"),
					),
					vecty.Class("button", "is-info", "is-small"),
					event.Click(s.onSimulateClick),
				),
				vecty.Text("Start simulation"),
			),
			vecty.If(s.hasResults,
				elem.Button(
					vecty.Markup(
						vecty.Class("button", "is-success", "is-small"),
						event.Click(s.onRestartClick),
					),
					vecty.Text("Replay"),
				),
				s.timeline,
			),
			elem.Break(),
			vecty.If(s.inProgress, elem.Div(
				vecty.Markup(
					vecty.Class("notification", "is-success"),
				),
				vecty.Text("Running simulation..."),
			)),
			elem.Div(
				vecty.If(s.errMsg != "", elem.Div(
					vecty.Markup(
						vecty.Class("notification", "is-danger"),
					),
					vecty.Text(s.errMsg),
				)),
			),
		),
	)
}

func (s *Simulation) onAddressChange(event *vecty.Event) {
	value := event.Target.Get("value").String()

	s.address = value
}

func (s *Simulation) onTTLChange(event *vecty.Event) {
	value := event.Target.Get("value").Int()

	s.ttl = value
}

// Address returns the current backend address.
func (s *Simulation) Address() string {
	return s.address
}

// TTL returns the current TTL value.
func (s *Simulation) TTL() int {
	return s.ttl
}

func (s *Simulation) onSimulateClick(e *vecty.Event) {
	go func() {
		s.errMsg = ""
		s.hasResults = false
		s.inProgress = true
		vecty.Rerender(s)

		steps, err := s.startSimulation()
		if err != nil {
			s.errMsg = err.Error()
		}

		s.hasResults = err == nil
		s.inProgress = false

		if s.hasResults {
			s.timeline = NewRange("Time", "", 0, 0, steps-1, s.step)
		}

		vecty.Rerender(s)
	}()
}

func (s *Simulation) onRestartClick(e *vecty.Event) {
	go s.replay()
}

func (s *Simulation) Reset() {
	s.hasResults = false
	s.inProgress = false
	s.errMsg = ""
	vecty.Rerender(s)
}

// StepForward increases timeline step.
func (s *Simulation) StepForward() {
	s.timeline.Inc()
}

// StepBackward decreases timeline step.
func (s *Simulation) StepBackward() {
	s.timeline.Dec()
}
