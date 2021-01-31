package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Button struct {
	state bool
	*Pane
	*Text
	rect *sdl.Rect
}

func (env *SDLEnvironment) NewButton(pne *Pane, txt *Text) *Button {
	var btn = &Button{false, pne, txt, &sdl.Rect{pne.x, pne.y, pne.w, pne.h}}
	btn.Text.x = btn.Pane.x + ((btn.Pane.w - btn.Text.w) / 2)
	btn.Text.y = btn.Pane.y + ((btn.Pane.h - btn.Text.h) / 2)
	return btn
}

func (btn *Button) render(r *sdl.Renderer) {
	btn.Pane.render(r)
	btn.Text.render(r)
}

func (btn *Button) renderActive(r *sdl.Renderer) {
	btn.Pane.renderActive(r)
	btn.Text.render(r)
}

func (btn *Button) inbounds(x, y int32) bool {
	return btn.Pane.inbounds(x, y)
}

func (btn *Button) active() *bool {
	return &btn.state
}

type Signature struct {
	state bool
	*HLine
	*Text
	rect *sdl.Rect
}

func (env *SDLEnvironment) NewSignature(hline *HLine, txt *Text) *Signature {
	hline.y = hline.y + txt.h
	txt.x = hline.x
	txt.y = hline.y - txt.h
	var s = &Signature{false, hline, txt, &sdl.Rect{hline.x, hline.y, hline.w, hline.thickness}}
	return s
}

func (s *Signature) render(r *sdl.Renderer) {
	s.HLine.render(r)
	s.Text.render(r)
}

func (s *Signature) renderActive(r *sdl.Renderer) {
	s.HLine.renderActive(r)
	s.Text.render(r)
}

func (sig *Signature) inbounds(x, y int32) bool {
	return sig.Text.inbounds(x, y)
}

func (sig *Signature) active() *bool {
	return &sig.state
}
