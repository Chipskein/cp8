package screen

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func Init() {
	log.Printf("Começou")
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Printf("%s\n", err)
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 200, 200, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Printf("%s\n", err)
		panic(err)
	}
	defer window.Destroy()

	//surface , err := window.GetSurface()
	//if err != nil {
	//	panic(err)
	//}
	//surface.FillRect(nil, 0)

	//rect := sdl.Rect{0, 0, 200, 200}
	//surface.FillRect(&rect, 0xffff0000)
	//window.UpdateSurface()
	/*
	   running := true

	   	for running {
	   		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	   			switch event.(type) {
	   			case *sdl.QuitEvent:
	   				println("Quit")
	   				running = false
	   				break
	   			}
	   		}
	   	}
	*/
}
