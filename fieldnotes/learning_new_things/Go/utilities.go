package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type color4 struct {
	r, g, b, a uint8
}

func (c *color4) SDLColor() sdl.Color {
	return sdl.Color{c.r, c.g, c.b, c.a}
}

type Point struct {
	x, y int32
}
