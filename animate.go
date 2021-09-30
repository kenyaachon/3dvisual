package main

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
	"github.com/hexops/vecty"
)

const (
	MagicNumberFPS = 60 // correct me if I'm wrong, but 60 FPS is a special value for 'requestAnimationFrame'
)

// animate fires up as an requestAnimationFrame handler.
func (w *WebGLScene) animate() {
	if w.renderer == nil {
		return
	}

	w.controls.Update()

	if w.fps == MagicNumberFPS {
		js.Global.Call("requestAnimationFrame", w.animate)
	} else {
		time.AfterFunc(time.Duration(1000/w.fps)*time.Millisecond, func() {
			js.Global.Call("requestAnimationFrame", w.animate)
		})
	}

	if w.autoRotate {
		pos := w.graphGroup.Object.Get("rotation")
		coeff := MagicNumberFPS / float64(w.fps) * 0.001 // rotate faster on lower FPS
		pos.Set("y", pos.Get("y").Float()+coeff)
		w.graphGroup.UpdateMatrix()
	}

	if w.wobble {
		w.wobbling.Animate()
		w.updatePositions()
	}

	// some render throttling magic to prevent wasting CPU/GPU while idle
	// if auto rotation or other effects are active, render always
	var needRendering bool = w.wobble || w.autoRotate
	if !needRendering {
		// else, consult render throttler
		needRendering = w.rt.NeedRendering()
	}

	if needRendering {
		w.renderer.Render(w.scene, w.camera)
		w.rt.ReenableIfNeeded()
	}
}

// ToggleAutoRotation switches auto rotation option.
func (w *WebGLScene) ToggleAutoRotation() {
	w.autoRotate = !w.autoRotate
}

// ToggleWobbling switches wobbling option.
func (w *WebGLScene) ToggleWobbling() {
	w.wobble = !w.wobble
}

// MouseMoveListener implements listener for mousemove events.
// We use it for disabling render throttling, as mousemove events
// correlates with user moving inside of the WebGL canvas. We
// may switch to use mousedown or drag events, but let's see how
// mousemove works.
// This is sort of a hack, as the proper approach would be to get
// data from controls code (w.controls.Update), but it's currently
// a JS code, so it's easier use this hack.
func (p *Page) MouseMoveListener(e *vecty.Event) {
	if !p.webgl.rt.NeedRendering() {
		p.webgl.rt.EnableRendering()
	}
}
