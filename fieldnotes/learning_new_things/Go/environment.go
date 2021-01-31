package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type SDLEnvironment struct {
	window          *sdl.Window
	renderer        *sdl.Renderer
	font            *ttf.Font
	fontSize        int8
	foregroundColor color4
	panelColor      color4
	panelEventColor color4
	textColor       color4
}

func NewSDLEnvironment(title string, width, height int32) *SDLEnvironment {
	se := new(SDLEnvironment)
	var err error
	err = sdl.Init(sdl.INIT_EVERYTHING)
	handle_panic(err, sdl.GetError())

	//Load default font and store
	err = ttf.Init()
	handle_errors(err, "initialize TTF")

	var window *sdl.Window
	window, err = sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	handle_errors(err, "create window")
	se.window = window

	var renderer *sdl.Renderer
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	handle_errors(err, "open renderer")
	se.renderer = renderer

	var font *ttf.Font
	font, err = ttf.OpenFont("/home/caseykneale/Desktop/GoGles/gofonts/Go-Regular.ttf", 12)
	handle_errors(err, "open font")
	se.font = font
	//Establish some defaults to make instantiating simple
	se.fontSize = 16
	se.foregroundColor = color4{110, 110, 110, 0}
	se.panelColor = color4{210, 210, 210, 0}
	se.panelEventColor = color4{160, 160, 160, 0}
	se.textColor = color4{0, 0, 0, 0}
	return se
}

func (env *SDLEnvironment) Render() {
	env.renderer.SetDrawColor(env.foregroundColor.r, env.foregroundColor.g, env.foregroundColor.b, env.foregroundColor.a)
	env.renderer.Clear()
}
