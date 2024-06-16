package cpu

import (
	"cp8/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/*UpdateDisplay function is used to set a display buffer to the screen.*/
func UpdateDisplay(DisplayBuffer [32][64]int) {
	for y_index, y := range DisplayBuffer {
		for x_index, x := range y {
			var color = rl.Black
			if x == 1 {
				color = rl.White
			}
			rl.DrawRectangle(config.Scale*config.PixelSize*int32(x_index), config.Scale*config.PixelSize*int32(y_index), config.PixelSize, config.PixelSize, color)
		}
	}
}
