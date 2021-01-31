package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func (s *Scene) foldChains(event sdl.Event) {
	var propagate bool
	for idx, chain := range s.actionChains {
		propagate = true
		for _, functor := range chain {
			if propagate {
				propagate = functor(s, s.actionIdx[idx], event)
			} else {
				break
			}
		}
	}
}

func (s *Scene) foldGlobalChains(event sdl.Event) {
	var propagate bool
	for _, chain := range s.globalChains {
		propagate = true
		for _, functor := range chain {
			if propagate {
				propagate = functor(s, event)
			} else {
				break
			}
		}
	}
}

func MouseButton(s *Scene, event sdl.Event) bool {
	var success = false
	switch event.(type) {
	case *sdl.MouseButtonEvent:
		success = true
	}
	return success
}

func ExitProgram(s *Scene, event sdl.Event) bool {
	switch event.(type) {
	case *sdl.QuitEvent:
		s.alive = false
	}
	return true
}

func DeactivateAll(s *Scene, event sdl.Event) bool {
	s.DeactivateAll()
	return true
}

func ValidIndex(s *Scene, idx int, event sdl.Event) bool {
	return idx > -1
}

func MouseDown(s *Scene, idx int, event sdl.Event) bool {
	var success = false
	switch e := event.(type) {
	case *sdl.MouseButtonEvent:
		switch e.GetType() {
		case sdl.MOUSEBUTTONDOWN:
			if (s.actionWidgets[idx]).inbounds(e.X, e.Y) {
				*(s.actionWidgets[idx]).active() = true
				success = true
			}
		}
	}
	return success
}

func MouseUp(s *Scene, idx int, event sdl.Event) bool {
	var success = false
	switch e := event.(type) {
	case *sdl.MouseButtonEvent:
		switch e.GetType() {
		case sdl.MOUSEBUTTONUP:
			if (s.actionWidgets[idx]).inbounds(e.X, e.Y) {
				*(s.actionWidgets[idx]).active() = false
				success = true
			}
		}
	}
	return success
}

func Activate(s *Scene, idx int, event sdl.Event) bool {
	*(s.actionWidgets[idx]).active() = true
	return true
}

func Deactivate(s *Scene, idx int, event sdl.Event) bool {
	*(s.actionWidgets[idx]).active() = false
	return true
}
