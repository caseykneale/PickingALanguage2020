package main

/*
Beware of he who would deny you access to information, for in his heart he dreams himself your master.
Commissioner Pravin Lal, "U.N. Declaration of Rights"
*/

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func customFunction(s *Scene, idx int, event sdl.Event) bool {
	fmt.Println("YOU PUSHED THE BUTTON!")
	return true
}

func main() {
	os.Setenv("DISPLAY", ":0")
	var winWidth, winHeight int32 = 800, 600
	var winTitle string = "Go-GUI"
	var mainWindow = NewSDLEnvironment(winTitle, winWidth, winHeight)

	var mainScene = NewScene()
	mainScene.AddGlobalChain(GlobalChain{ExitProgram})
	mainScene.AddGlobalChain(GlobalChain{MouseButton, DeactivateAll})

	var btnTxt = NewText(mainWindow, 0, 0, "Submit")
	var btnFace = NewPane(mainWindow, Point{32, 32}, 64, 32)
	var newp = mainWindow.NewButton(btnFace, btnTxt)
	mainScene.AddActive(newp)

	mainScene.AddChain(WidgetChain{MouseDown})
	mainScene.AddChain(WidgetChain{MouseUp, customFunction})

	var sigTxt = NewText(mainWindow, 0, 0, "Enter Text Here")
	var sigLine = NewHLine(mainWindow, Point{128, 32}, 96, 2)
	var newt = mainWindow.NewSignature(sigLine, sigTxt)
	mainScene.AddActive(newt)

	var frames = 0
	var start = time.Now()

	for mainScene.alive {
		//begin rendering foreground
		mainWindow.renderer.Present()
		mainWindow.Render()
		//Handle widget events
		mainScene.EventHandler()
		//Draw widgets
		mainScene.render(mainWindow.renderer)

		frames++
		var t = time.Now()
		var elapsed = t.Sub(start)
		if frames%1000 == 0 {
			fmt.Printf("FPS %d\n", 1000.0/(elapsed.Seconds()))
			start = time.Now()
		}
	}

	defer sdl.Quit()
	defer mainWindow.window.Destroy()
	mainWindow.font.Close() //add to destructor
	mainWindow.renderer.Destroy()
	mainWindow.window.Destroy()
	sdl.Quit()
}
