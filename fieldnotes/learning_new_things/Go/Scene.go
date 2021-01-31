package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type WidgetChain = []func(s *Scene, idx int, event sdl.Event) bool
type GlobalChain = []func(s *Scene, event sdl.Event) bool

type Scene struct {
	alive             bool
	actionWidgets     []ActiveWidget
	decorationWidgets []Widget
	actionIdx         []int
	actionChains      []WidgetChain
	globalChains      []GlobalChain
}

func NewScene() *Scene {
	return &Scene{
		true,
		make([]ActiveWidget, 0),
		make([]Widget, 0),
		make([]int, 0),
		make([][]func(s *Scene, idx int, event sdl.Event) bool, 0),
		make([][]func(s *Scene, event sdl.Event) bool, 0),
	}
}

func (s *Scene) render(r *sdl.Renderer) {
	for _, widget := range s.decorationWidgets {
		widget.render(r)
	}
	for _, widget := range s.actionWidgets {
		if *widget.active() == true {
			widget.renderActive(r)
		} else {
			widget.render(r)
		}
	}
}

func (s *Scene) AddDecor(decoration Widget) {
	s.decorationWidgets = append(s.decorationWidgets, decoration)
}

func (s *Scene) AddActive(action ActiveWidget) {
	s.actionWidgets = append(s.actionWidgets, action)
}

func (s *Scene) AddChain(fnchain []func(s *Scene, idx int, event sdl.Event) bool) {
	s.actionChains = append(s.actionChains, fnchain)
	s.actionIdx = append(s.actionIdx, len(s.actionWidgets)-1)
}

func (s *Scene) AddGlobalChain(fnchain []func(s *Scene, event sdl.Event) bool) {
	s.globalChains = append(s.globalChains, fnchain)
}

func (s *Scene) EventHandler() bool {
	var event sdl.Event
	var running = true
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		s.foldGlobalChains(event)
		s.foldChains(event)
	}
	return running
}

func (s *Scene) DeactivateAll() {
	for _, widget := range s.actionWidgets {
		*widget.active() = false
	}
}
