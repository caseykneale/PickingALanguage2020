package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Widget interface {
	render(r *sdl.Renderer)
}

type ActiveWidget interface {
	render(r *sdl.Renderer)
	renderActive(r *sdl.Renderer)
	inbounds(x, y int32) bool
	active() *bool
}

type HLine struct {
	state bool
	Point
	w         int32
	thickness int32
	color     color4
}

func NewHLine(env *SDLEnvironment, pos Point, w, thickness int32) *HLine {
	return &HLine{false, pos, w, thickness, env.textColor}
}

func (hline *HLine) render(r *sdl.Renderer) {
	var rect = sdl.Rect{hline.x, hline.y, hline.w, hline.thickness}
	r.SetDrawColor(hline.color.r, hline.color.g, hline.color.b, hline.color.a)
	r.FillRect(&rect)
}

func (hline *HLine) renderActive(r *sdl.Renderer) {
	hline.render(r)
}

func (hline *HLine) inbounds(x, y int32) bool {
	return (x > hline.x) && (x < hline.x+hline.w) && (y > hline.y) && (x > hline.y+hline.thickness)
}

func (hline *HLine) active() *bool {
	return &hline.state
}

type Text struct {
	state bool
	Point
	font     *ttf.Font
	fontSize int8
	w        int32
	h        int32
	color    color4
	text     string
	texture  *sdl.Texture
}

func NewText(env *SDLEnvironment, x, y int32, text string) *Text {
	var err error
	var blendedSurface *sdl.Surface
	blendedSurface, err = env.font.RenderUTF8Blended(text, env.textColor.SDLColor())
	handle_errors(err, "render text")

	var w = blendedSurface.W
	var h = blendedSurface.H

	var blendedTexture *sdl.Texture
	blendedTexture, err = env.renderer.CreateTextureFromSurface(blendedSurface)
	handle_errors(err, "create text texture")
	blendedSurface.Free()

	return &Text{false, Point{x, y}, env.font, env.fontSize, w, h, env.textColor, text, blendedTexture}
}

func (txt *Text) render(r *sdl.Renderer) {
	var txtRect = sdl.Rect{txt.x, txt.y, txt.w, txt.h}
	r.SetDrawColor(txt.color.r, txt.color.g, txt.color.b, txt.color.a)
	r.Copy(txt.texture, nil, &txtRect)
}

func (txt *Text) renderActive(r *sdl.Renderer) {
	txt.render(r)
}

func (txt *Text) inbounds(x, y int32) bool {
	return (x > txt.x) && (x < txt.x+txt.w) && (y > txt.y) && (x > txt.y+txt.h)
}

func (txt *Text) active() *bool {
	return &txt.state
}

type Pane struct {
	state bool
	Point
	w          int32
	h          int32
	color      color4
	eventColor color4
}

func NewPane(env *SDLEnvironment, pos Point, w, h int32) *Pane {
	return &Pane{false, pos, w, h, env.panelColor, env.panelEventColor}
}

func (pne *Pane) render(r *sdl.Renderer) {
	var rect = sdl.Rect{pne.x, pne.y, pne.w, pne.h}
	r.SetDrawColor(pne.color.r, pne.color.g, pne.color.b, pne.color.a)
	r.FillRect(&rect)
}

func (pne *Pane) renderActive(r *sdl.Renderer) {
	var rect = sdl.Rect{pne.x, pne.y, pne.w, pne.h}
	r.SetDrawColor(pne.eventColor.r, pne.eventColor.g, pne.eventColor.b, pne.eventColor.a)
	r.FillRect(&rect)
}

func (pne *Pane) inbounds(x, y int32) bool {
	return (x > pne.x) && (x < pne.x+pne.w) && (y > pne.y) && (x > pne.y+pne.h)
}

func (pne *Pane) active() *bool {
	return &pne.state
}
